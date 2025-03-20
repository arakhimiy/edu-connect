package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2689671926")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_4192607770",
			"hidden": false,
			"id": "relation2063623452",
			"maxSelect": 1,
			"minSelect": 0,
			"name": "status",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_2689671926")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("relation2063623452")

		return app.Save(collection)
	})
}
