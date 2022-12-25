package migration

import (
	"backend/db/migration/menus"
	"backend/db/migration/users"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const DATABASE_NAME = "test"

func Migration(client *mongo.Client) {

	db := client.Database(DATABASE_NAME)

	users.CreateUsers(db)

	menus.CreateMenus(db)

}

func CollectionDelete(client *mongo.Client) {

	db := client.Database(DATABASE_NAME)

	users.DropUsers(db)

	menus.DropMenus(db)

}

func DatabaseDrop(client *mongo.Client) {

	database := client.Database(DATABASE_NAME)

	err := database.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}
