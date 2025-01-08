package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2125759099")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(2, []byte(`{
			"hidden": false,
			"id": "select2812505443",
			"maxSelect": 1,
			"name": "degree",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "select",
			"values": [
				"bachelor",
				"master"
			]
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2125759099")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("select2812505443")

		return app.Save(collection)
	})
}
