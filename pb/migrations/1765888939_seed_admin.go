package migrations

import (
	"errors"
	"os"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		email := os.Getenv("PB_EMAIL")
		password := os.Getenv("PB_PASSWORD")
		if email == "" || password == "" {
			return errors.New("PB_EMAIL and PB_PASSWORD must be set")
		}
		
		col, _ := app.FindCollectionByNameOrId(core.CollectionNameSuperusers)
		rec := core.NewRecord(col)
		rec.Set("email", email)
		rec.Set("password", password)
		return app.Save(rec)
	}, func(app core.App) error {
		email := os.Getenv("PB_EMAIL")
		rec, _ := app.FindAuthRecordByEmail(core.CollectionNameSuperusers, email)
		if rec == nil {
			return nil
		}
		return app.Delete(rec)

	})
}
