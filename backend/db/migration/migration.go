package migration

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Tea struct {
	Type     string
	Category string
	Toppings []string
	Price    float32
}

func Migration(client *mongo.Client) {

	coll := client.Database("tea").Collection("menu")
	docs := []interface{}{
		Tea{Type: "Masala", Category: "black", Toppings: []string{"ginger", "pumpkin spice", "cinnamon"}, Price: 6.75},
		Tea{Type: "Gyokuro", Category: "green", Toppings: []string{"berries", "milk foam"}, Price: 5.65},
		Tea{Type: "English Breakfast", Category: "black", Toppings: []string{"whipped cream", "honey"}, Price: 5.75},
		Tea{Type: "Sencha", Category: "green", Toppings: []string{"lemon", "whipped cream"}, Price: 5.15},
		Tea{Type: "Assam", Category: "black", Toppings: []string{"milk foam", "honey", "berries"}, Price: 5.65},
		Tea{Type: "Matcha", Category: "green", Toppings: []string{"whipped cream", "honey"}, Price: 6.45},
		Tea{Type: "Earl Grey", Category: "black", Toppings: []string{"milk foam", "pumpkin spice"}, Price: 6.15},
		Tea{Type: "Hojicha", Category: "green", Toppings: []string{"lemon", "ginger", "milk foam"}, Price: 5.55},
	}
	result, err := coll.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Print(result)

}

func CollectionDrop(client *mongo.Client) {

	coll := client.Database("tea").Collection("menu")

	err := coll.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}

func DatabaseDrop(client *mongo.Client) {

	database := client.Database("tea")

	err := database.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}
