package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(18, []byte(`{
			"hidden": false,
			"id": "select1212835736",
			"maxSelect": 1,
			"name": "maritialStatus",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"married",
				"alone"
			]
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(19, []byte(`{
			"hidden": false,
			"id": "select3343321666",
			"maxSelect": 1,
			"name": "gender",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"male",
				"female"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("_pb_users_auth_")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("select1212835736")

		// remove field
		collection.Fields.RemoveById("select3343321666")

		return app.Save(collection)
	})
}
