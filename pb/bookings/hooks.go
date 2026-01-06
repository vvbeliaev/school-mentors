package bookings

import (
	"errors"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func Hooks(app *pocketbase.PocketBase) {
	// Atomically book the slot and set booking metadata
	app.OnRecordCreate("bookings").BindFunc(func(e *core.RecordEvent) error {
		slotId := e.Record.GetString("slot")

		// Atomic UPDATE: only succeeds if slot exists AND is not yet booked
		// This prevents race conditions (two users booking the same slot simultaneously)
		res, err := e.App.DB().NewQuery("UPDATE slots SET booked = 1 WHERE id = {:id} AND (booked = 0 OR booked IS NULL)").
			Bind(map[string]any{"id": slotId}).
			Execute()
		if err != nil {
			return err
		}

		rows, _ := res.RowsAffected()
		if rows == 0 {
			return errors.New("slot is already booked or does not exist")
		}

		// Fetch slot to get mentor info for pricing
		slot, err := e.App.FindRecordById("slots", slotId)
		if err != nil {
			return err
		}

		// Copy mentor's current hourly rate
		mentorId := slot.GetString("mentor")
		mentor, err := e.App.FindRecordById("users", mentorId)
		if err == nil {
			e.Record.Set("agreedPriceCents", mentor.GetInt("hourRateCents"))
		}

		e.Record.Set("status", "pending")

		return e.Next()
	})

	// When a booking is canceled or expired -> free the slot
	app.OnRecordUpdate("bookings").BindFunc(func(e *core.RecordEvent) error {
		status := e.Record.GetString("status")
		if status == "canceled" || status == "expired" {
			slotId := e.Record.GetString("slot")
			slot, err := app.FindRecordById("slots", slotId)
			if err == nil {
				slot.Set("booked", false)
				app.Save(slot)
			}
		}
		return e.Next()
	})
}
