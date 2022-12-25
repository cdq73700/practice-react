package menus

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Menus struct {
	Type     string
	Category string
	Toppings []string
	Price    float32
}

const MENU = "menu"

func CreateMenus(db *mongo.Database) {
	menu := db.Collection(MENU)
	docs := []interface{}{
		&Menus{Type: "Masala", Category: "black", Toppings: []string{"ginger", "pumpkin spice", "cinnamon"}, Price: 6.75},
		&Menus{Type: "Gyokuro", Category: "green", Toppings: []string{"berries", "milk foam"}, Price: 5.65},
		&Menus{Type: "English Breakfast", Category: "black", Toppings: []string{"whipped cream", "honey"}, Price: 5.75},
		&Menus{Type: "Sencha", Category: "green", Toppings: []string{"lemon", "whipped cream"}, Price: 5.15},
		&Menus{Type: "Assam", Category: "black", Toppings: []string{"milk foam", "honey", "berries"}, Price: 5.65},
		&Menus{Type: "Matcha", Category: "green", Toppings: []string{"whipped cream", "honey"}, Price: 6.45},
		&Menus{Type: "Earl Grey", Category: "black", Toppings: []string{"milk foam", "pumpkin spice"}, Price: 6.15},
		&Menus{Type: "Hojicha", Category: "green", Toppings: []string{"lemon", "ginger", "milk foam"}, Price: 5.55},
	}
	result, err := menu.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Print(result)
}

func DropMenus(db *mongo.Database) {

	menu := db.Collection(MENU)

	err := menu.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}
