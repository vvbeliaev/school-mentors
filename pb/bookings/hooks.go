package bookings

import (
	"fmt"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func Hooks(app *pocketbase.PocketBase) {
	// Check if the slot is already booked
	// When a booking is created -> book the slot
	// Copy the mentor's hourly rate to the booking in case the mentor changes their rate
	app.OnRecordCreate("bookings").BindFunc(func(e *core.RecordEvent) error {
		slotId := e.Record.GetString("slot")
		slot, err := app.FindRecordById("slots", slotId)
		if err != nil {
			return fmt.Errorf("slot not found: %w", err)
		}

		if slot.GetBool("isBooked") {
			return fmt.Errorf("slot is already booked: %s", slotId)
		}
		
		slot.Set("isBooked", true)
		if err := app.Save(slot); err != nil {
			return err
		}

		mentorId := slot.GetString("mentor")
		mentor, err := app.FindRecordById("users", mentorId)
		if err == nil {
			currentRate := mentor.GetInt("hourRateCents")
			e.Record.Set("agreedPriceCents", currentRate)
		}

		e.Record.Set("status", "pending")

		return e.Next()
	})

	// When a booking is canceled -> free the slot
	app.OnRecordUpdate("bookings").BindFunc(func(e *core.RecordEvent) error {
		status := e.Record.GetString("status")
		if status == "canceled" || status == "expired" {
			slotId := e.Record.GetString("slot")
			slot, err := app.FindRecordById("slots", slotId)
			if err == nil {
				slot.Set("isBooked", false)
				app.Save(slot)
			}
		}
		return e.Next()
	})
}
