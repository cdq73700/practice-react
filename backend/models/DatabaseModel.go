package models

import "go.mongodb.org/mongo-driver/mongo"

type ConnectionDatabase struct {
	*mongo.Collection
}