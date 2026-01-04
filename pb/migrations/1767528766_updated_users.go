package migrations

import (
	"encoding/json"

	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"deleteRule": "id = @request.auth.id\n&&\n@request.body.isVerified:isset = false\n&&\n@request.body.stripeCustomer:isset = false",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_tokenKey__pb_users_auth_` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `tokenKey` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_email__pb_users_auth_` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `email` + "`" + `) WHERE ` + "`" + `email` + "`" + ` != ''",
				"CREATE INDEX ` + "`" + `idx_fhFPXaESmr` + "`" + ` ON ` + "`" + `users` + "`" + ` (\n  ` + "`" + `isMentor` + "`" + `,\n  ` + "`" + `isVerified` + "`" + `\n)",
				"CREATE INDEX ` + "`" + `idx_BwBHmRGySQ` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `university` + "`" + `)",
				"CREATE INDEX ` + "`" + `idx_efGRmUSRPa` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `stripeCustomer` + "`" + `)"
			],
			"oauth2": {
				"enabled": true
			},
			"updateRule": "id = @request.auth.id\n&&\n@request.body.isVerified:isset = false\n&&\n@request.body.stripeCustomer:isset = false"
		}`), &collection); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(8, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text2692384236",
			"max": 0,
			"min": 0,
			"name": "university",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(9, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text3709889147",
			"max": 0,
			"min": 0,
			"name": "bio",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(10, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text2812505443",
			"max": 0,
			"min": 0,
			"name": "degree",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(11, []byte(`{
			"hidden": false,
			"id": "number971074285",
			"max": null,
			"min": 1950,
			"name": "graduationYear",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(12, []byte(`{
			"hidden": false,
			"id": "json1874629670",
			"maxSize": 0,
			"name": "tags",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "json"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(13, []byte(`{
			"hidden": false,
			"id": "bool1098737668",
			"name": "isMentor",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "bool"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(14, []byte(`{
			"hidden": false,
			"id": "bool2487745274",
			"name": "isVerified",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "bool"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(15, []byte(`{
			"autogeneratePattern": "",
			"hidden": false,
			"id": "text599801048",
			"max": 0,
			"min": 0,
			"name": "stripeCustomer",
			"pattern": "",
			"presentable": false,
			"primaryKey": false,
			"required": false,
			"system": false,
			"type": "text"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(16, []byte(`{
			"hidden": false,
			"id": "number2949507886",
			"max": null,
			"min": 0,
			"name": "hourRateCents",
			"onlyInt": true,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// update collection data
		if err := json.Unmarshal([]byte(`{
			"deleteRule": "id = @request.auth.id",
			"indexes": [
				"CREATE UNIQUE INDEX ` + "`" + `idx_tokenKey__pb_users_auth_` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `tokenKey` + "`" + `)",
				"CREATE UNIQUE INDEX ` + "`" + `idx_email__pb_users_auth_` + "`" + ` ON ` + "`" + `users` + "`" + ` (` + "`" + `email` + "`" + `) WHERE ` + "`" + `email` + "`" + ` != ''"
			],
			"oauth2": {
				"enabled": false
			},
			"updateRule": "id = @request.auth.id"
		}`), &collection); err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("text2692384236")

		// remove field
		collection.Fields.RemoveById("text3709889147")

		// remove field
		collection.Fields.RemoveById("text2812505443")

		// remove field
		collection.Fields.RemoveById("number971074285")

		// remove field
		collection.Fields.RemoveById("json1874629670")

		// remove field
		collection.Fields.RemoveById("bool1098737668")

		// remove field
		collection.Fields.RemoveById("bool2487745274")

		// remove field
		collection.Fields.RemoveById("text599801048")

		// remove field
		collection.Fields.RemoveById("number2949507886")

		return app.Save(collection)
	})
}
