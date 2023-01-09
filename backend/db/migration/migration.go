package migration

import (
	model "backend/Model"
	"backend/db/migration/menus"
	"backend/db/migration/users"
	"context"
)

func Migration() {

	users.CreateUsers()

	menus.CreateMenus()

}

func CollectionDelete() {

	users.DropUsers()

	menus.DropMenus()

}

func DatabaseDrop() {

	err := model.DataBaseClient.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}
