package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://admin:admin@fiber-mongo:27017"
const DataBaseName = "test"

var Client *mongo.Client
var DataBaseClient *mongo.Database

var Err error

func Initialized() {
	GetMongoDbConnection()
	GetMongoDataBase()
	UserInitialized()
	MenuInitialized()
}

func GetMongoDbConnection() {

	//Connect to MongoDB
	Client, Err = mongo.NewClient(options.Client().ApplyURI(uri))
	if Err != nil {
		log.Fatal(Err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	Err = Client.Connect(ctx)
	if Err != nil {
		log.Fatal(Err)
	}

	fmt.Println("connection success:")
}

func GetMongoDataBase() {

	DataBaseClient = Client.Database(DataBaseName)

}

func GetMongoDbCollection(CollectionName string) *mongo.Collection {

	return DataBaseClient.Collection(CollectionName)

}
