package menus

import (
	model "backend/model"
	"context"
	"fmt"
)


func CreateMenus() {
	menu := model.MenuCollectionClient
	docs := []interface{}{
		&model.MenuStruct{Type: "Masala", Category: "black", Toppings: []string{"ginger", "pumpkin spice", "cinnamon"}, Price: 6.75},
		&model.MenuStruct{Type: "Gyokuro", Category: "green", Toppings: []string{"berries", "milk foam"}, Price: 5.65},
		&model.MenuStruct{Type: "English Breakfast", Category: "black", Toppings: []string{"whipped cream", "honey"}, Price: 5.75},
		&model.MenuStruct{Type: "Sencha", Category: "green", Toppings: []string{"lemon", "whipped cream"}, Price: 5.15},
		&model.MenuStruct{Type: "Assam", Category: "black", Toppings: []string{"milk foam", "honey", "berries"}, Price: 5.65},
		&model.MenuStruct{Type: "Matcha", Category: "green", Toppings: []string{"whipped cream", "honey"}, Price: 6.45},
		&model.MenuStruct{Type: "Earl Grey", Category: "black", Toppings: []string{"milk foam", "pumpkin spice"}, Price: 6.15},
		&model.MenuStruct{Type: "Hojicha", Category: "green", Toppings: []string{"lemon", "ginger", "milk foam"}, Price: 5.55},
	}
	result, err := menu.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Print(result)
}

func DropMenus() {

	menu := model.MenuCollectionClient

	err := menu.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}
