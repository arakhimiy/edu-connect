package migrations

import (
	"github.com/pocketbase/pocketbase/core"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3036160472")
		if err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(4, []byte(`{
			"hidden": false,
			"id": "number3485090271",
			"max": null,
			"min": null,
			"name": "tuitionFee",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(5, []byte(`{
			"hidden": false,
			"id": "number60950023",
			"max": null,
			"min": null,
			"name": "courseDuration",
			"onlyInt": false,
			"presentable": false,
			"required": false,
			"system": false,
			"type": "number"
		}`)); err != nil {
			return err
		}

		// add field
		if err := collection.Fields.AddMarshaledJSONAt(6, []byte(`{
			"cascadeDelete": false,
			"collectionId": "pbc_2125759099",
			"hidden": false,
			"id": "relation446653667",
			"maxSelect": 999,
			"minSelect": 0,
			"name": "faculties",
			"presentable": false,
			"required": false,
			"system": false,
			"type": "relation"
		}`)); err != nil {
			return err
		}

		return app.Save(collection)
	}, func(app core.App) error {
		collection, err := app.FindCollectionByNameOrId("pbc_3036160472")
		if err != nil {
			return err
		}

		// remove field
		collection.Fields.RemoveById("number3485090271")

		// remove field
		collection.Fields.RemoveById("number60950023")

		// remove field
		collection.Fields.RemoveById("relation446653667")

		return app.Save(collection)
	})
}
