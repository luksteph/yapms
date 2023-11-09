package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models/schema"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("0v1w5g3emykgsnt")
		if err != nil {
			return err
		}

		// update
		edit_enabled := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "voy4kpum",
			"name": "enabled",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_enabled)
		collection.Schema.AddField(edit_enabled)

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("0v1w5g3emykgsnt")
		if err != nil {
			return err
		}

		// update
		edit_enabled := &schema.SchemaField{}
		json.Unmarshal([]byte(`{
			"system": false,
			"id": "voy4kpum",
			"name": "approved",
			"type": "bool",
			"required": false,
			"presentable": false,
			"unique": false,
			"options": {}
		}`), edit_enabled)
		collection.Schema.AddField(edit_enabled)

		return dao.SaveCollection(collection)
	})
}
