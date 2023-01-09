package migration

import (
	"backend/db/migration/menus"
	"backend/db/migration/users"
	model "backend/model"
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
