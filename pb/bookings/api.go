package bookings

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/checkout/session"
	"github.com/stripe/stripe-go/v84/customer"
	"github.com/stripe/stripe-go/v84/webhook"
)

func API(se *core.ServeEvent, app *pocketbase.PocketBase) {
	se.Router.POST("/api/bookings/checkout/{id}", func(e *core.RequestEvent) error {
		if e.Auth == nil {
			return e.UnauthorizedError("You must be logged in to pay", nil)
		}

		bookingId := e.Request.PathValue("id")

		booking, err := app.FindRecordById("bookings", bookingId)
		if err != nil {
			return e.NotFoundError("Booking not found", err)
		}

		if booking.GetString("mentee") != e.Auth.Id {
			return e.ForbiddenError("You are not allowed to pay for this booking", nil)
		}

		if booking.GetString("status") != "pending" {
			return e.BadRequestError("Only pending bookings can be paid", nil)
		}

		slotId := booking.GetString("slot")
		slot, err := app.FindRecordById("slots", slotId)
		if err != nil {
			return e.InternalServerError("Slot not found", err)
		}

		mentorId := slot.GetString("mentor")
		mentor, err := app.FindRecordById("users", mentorId)
		if err != nil {
			return e.InternalServerError("Mentor not found", err)
		}

		price := int64(booking.GetInt("agreedPriceCents"))
		if price == 0 {
			return e.BadRequestError("Price is not set", nil)
		}

		domain := os.Getenv("PUBLIC_DOMAIN")
		if domain == "" {
			return e.InternalServerError("PUBLIC_DOMAIN is not set", nil)
		}

		stripeCustID := e.Auth.GetString("stripeCustomer")
		if stripeCustID == "" {
			custParams := &stripe.CustomerParams{
				Email: stripe.String(e.Auth.GetString("email")),
				Name:  stripe.String(e.Auth.GetString("name")),
				Metadata: map[string]string{
					"userId": e.Auth.Id,
				},
			}
			cust, err := customer.New(custParams)
			if err != nil {
				return e.InternalServerError("Failed to create stripe customer", err)
			}
			stripeCustID = cust.ID

			e.Auth.Set("stripeCustomer", stripeCustID)
			if err := app.Save(e.Auth); err != nil {
				return e.InternalServerError("Failed to update user with stripe ID", err)
			}
		}

		params := &stripe.CheckoutSessionParams{
			Customer:   stripe.String(stripeCustID),
			SuccessURL: stripe.String(domain + "/app/bookings/success?bookingId=" + bookingId),
			CancelURL:  stripe.String(domain + "/app/bookings/cancel?bookingId=" + bookingId),
			PaymentMethodTypes: stripe.StringSlice([]string{
				"card",
			}),
			Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
						Currency: stripe.String("usd"),
						ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
							Name: stripe.String("Mentorship Session with " + mentor.GetString("name")),
							Description: stripe.String("Duration: " + slot.GetString("durationMinutes") + " min"),
						},
						UnitAmount: stripe.Int64(price),
					},
					Quantity: stripe.Int64(1),
				},
			},
			Metadata: map[string]string{
				"bookingId": bookingId,
			},

			// TODO: Enable invoice creation
			// InvoiceCreation: &stripe.CheckoutSessionInvoiceCreationParams{
			// 	Enabled: stripe.Bool(true),
			// },
		}

		s, err := session.New(params)
		if err != nil {
			return e.InternalServerError("Failed to create checkout session", err)
		}

		return e.JSON(http.StatusOK, map[string]string{
			"url": s.URL,
		})
	})

	// 2. Webhook Handler
	se.Router.POST("/api/bookings/webhook", func(e *core.RequestEvent) error {
		const MaxBodyBytes = int64(65536)
		e.Request.Body = http.MaxBytesReader(e.Response, e.Request.Body, MaxBodyBytes)
		payload, err := io.ReadAll(e.Request.Body)
		if err != nil {
			return e.BadRequestError("Error reading request body", err)
		}

		sigHeader := e.Request.Header.Get("Stripe-Signature")
		endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

		event, err := webhook.ConstructEvent(payload, sigHeader, endpointSecret)
		if err != nil {
			return e.BadRequestError("Invalid signature", err)
		}
		
		if event.Type == "checkout.session.completed" {
			var sessionStruct stripe.CheckoutSession
			err := json.Unmarshal(event.Data.Raw, &sessionStruct)
			if err != nil {
				return e.BadRequestError("Error unmarshaling session", err)
			}

			bookingId := sessionStruct.Metadata["bookingId"]
			if bookingId == "" {
				return e.NoContent(http.StatusOK) 
			}

			booking, err := app.FindRecordById("bookings", bookingId)
			if err != nil {
				app.Logger().Error("Webhook: Booking not found", "bookingId", bookingId)
				return e.NoContent(http.StatusOK)
			}

			if booking.GetString("status") == "confirmed" {
				return e.NoContent(http.StatusOK)
			}

			booking.Set("status", "confirmed")
			
			if sessionStruct.PaymentIntent != nil {
				booking.Set("stripePaymentIntent", sessionStruct.PaymentIntent.ID)
			}

			if err := app.Save(booking); err != nil {
				app.Logger().Error("Webhook: Failed to save booking", "err", err)
				return e.InternalServerError("Failed to update booking status", err)
			}
		}

		return e.NoContent(http.StatusOK)
	})
}