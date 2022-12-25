package users

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Users struct {
	Email    string
	Username string
	Password string
}

const USERS = "users"

func CreateUsers(db *mongo.Database) {
	users := db.Collection(USERS)
	docs := []interface{}{
		&Users{
			Email:    "test@test.com",
			Username: "test",
			Password: "test",
		},
		&Users{
			Email:    "test2@test.com",
			Username: "test2",
			Password: "test2"},
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
