package users

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	Id				primitive.ObjectID	"_id"
	Email			string
	Name			string
	Password	string
}

const USERS = "users"

func CreateUsers(db *mongo.Database) {
	users := db.Collection(USERS)
	docs := []interface{}{
		&Users{
			Id:				primitive.NewObjectID(),
			Email:		"test@test.com",
			Name:			"test",
			Password:	"test",
		},
		&Users{
			Id:				primitive.NewObjectID(),
			Email:		"test2@test.com",
			Name:			"test2",
			Password:	"test2",},
	}

	result, err := users.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Print(result)
}

func DropUsers(db *mongo.Database) {

	users := db.Collection(USERS)

	err := users.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}
