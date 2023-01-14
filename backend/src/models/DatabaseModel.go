package models

import "go.mongodb.org/mongo-driver/mongo"


type MongoStrict struct {
	DbName string
	CollName string
	Client *mongo.Client
	Database *mongo.Database
	Collection *mongo.Collection
}

type MongoInterface interface {
	ConnectionDatabase() error
	SetDatabase(string)
	SetCollection(string)
}