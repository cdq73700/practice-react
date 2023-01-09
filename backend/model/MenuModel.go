package model

import "go.mongodb.org/mongo-driver/mongo"

const MENU = "menu"

var MenuCollectionClient *mongo.Collection

type MenuStruct struct {
	Type     string
	Category string
	Toppings []string
	Price    float32
}

func MenuInitialized () {
	MenuCollectionClient = GetMongoDbCollection(MENU)
}
