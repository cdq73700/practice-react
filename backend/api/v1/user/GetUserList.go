package v1

import (
	"backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userStruct models.UserStruct

func GetUserList(coll *mongo.Collection) ([]userStruct,error) {
	// Initialize variables.
	var users []userStruct

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	res, err := coll.Find(customContext, bson.D{})
	if err != nil {
		return nil, err
	}
	defer res.Close(customContext)

	if err = res.All(customContext, users); err != nil {
		return nil, err
	}

	// Return all of our users.
	// Return results.
	return users, nil
}