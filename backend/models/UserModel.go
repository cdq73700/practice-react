package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStruct struct {
	Id primitive.ObjectID `bson:"_id"`
	Email string `bson:"email"`
	Name string `bson:"name"`
	Password []byte `bson:"password"`
} 
