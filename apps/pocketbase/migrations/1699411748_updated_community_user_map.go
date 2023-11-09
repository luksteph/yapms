package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("34af5seyuzjymeg")
		if err != nil {
			return err
		}

		collection.Name = "community_user_maps"

		return dao.SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("34af5seyuzjymeg")
		if err != nil {
			return err
		}

		collection.Name = "community_user_map"

		return dao.SaveCollection(collection)
	})
}
