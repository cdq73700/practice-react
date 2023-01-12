package database

import (
	"backend/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://admin:admin@fiber-mongo:27017"

type MongoDbCollection models.ConnectionDatabase

// GetMongoDbConnection get connection of mongodb
func getMongoDbConnection() (*mongo.Client, error) {

	//Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection success:")

	return client, nil
}

func GetMongoDbCollection(DbName string, CollectionName string) (*MongoDbCollection, error) {
	client, err := getMongoDbConnection()

	if err != nil {
		return nil, err
	}

	collection := client.Database(DbName).Collection(CollectionName)

	return &MongoDbCollection{Collection: collection}, nil
}