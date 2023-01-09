package users

import (
	model "backend/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUsers() {
	users := model.UserCollectionClient
	docs := []interface{}{
		&model.UserStruct{
			Id:				primitive.NewObjectID(),
			Email:		"test@test.com",
			Name:			"test",
			Password:	model.CreatePassword("test"),
		},
		&model.UserStruct{
			Id:				primitive.NewObjectID(),
			Email:		"test2@test.com",
			Name:			"test2",
			Password:	model.CreatePassword("test"),},
	}

	result, err := users.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	fmt.Print(result)
}

func DropUsers() {

	users := model.UserCollectionClient

	err := users.Drop(context.TODO())

	if err != nil {
		panic(err)
	}

}