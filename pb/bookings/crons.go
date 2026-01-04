package bookings

import (
	"fmt"
	"time"

	"github.com/pocketbase/pocketbase"
)

func Crons(app *pocketbase.PocketBase) {
	app.Cron().MustAdd("expireOldBookings", "*/10 * * * *", func() {
		collection, err := app.FindCollectionByNameOrId("bookings")
		if err != nil {
			fmt.Println("Could not find bookings collection:", err)
			return
		}

		cutoffTime := time.Now().UTC().Add(-15 * time.Minute)
		cutoffStr := cutoffTime.Format("2006-01-02 15:04:05.000Z")
		
		records, err := app.FindRecordsByFilter(collection, fmt.Sprintf("status = 'pending' AND created < '%s'", cutoffStr), "", 0, 0)
		if err != nil {
			fmt.Println("Error querying old bookings:", err)
			return
		}

		for _, rec := range records {
			rec.Set("status", "expired")
			if err := app.Save(rec); err != nil {
				fmt.Printf("Failed to expire booking %s: %v\n", rec.Id, err)
			}
		}
	})
}