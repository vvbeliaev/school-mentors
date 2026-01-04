package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_53644091")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_4mgmbOzynI` + "`" + ` ON ` + "`" + `slots` + "`" + ` (` + "`" + `mentor` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_TclQdwl2EY` + "`" + ` ON ` + "`" + `slots` + "`" + ` (` + "`" + `booked` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"hidden": false,
			"id": "bool2584614489",
			"name": "booked",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "bool"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_53644091")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"indexes": [
				"CREATE INDEX ` + "`" + `idx_4mgmbOzynI` + "`" + ` ON ` + "`" + `slots` + "`" + ` (` + "`" + `mentor` + "`" + `)"
			]
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("bool2584614489")

		return app.Save(collection)
	})
}
