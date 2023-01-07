package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserUpdate(c *fiber.Ctx, db *mongo.Database) error {
	userDb := db.Collection(USERS)

	data := new(UserStruct)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: data.Id}}
	
	docs := &UserStruct {
		Id:				data.Id,
		Email:		data.Email,
		Name:			data.Name,
		Password:	data.Password,
	}

	result, err := userDb.UpdateOne(context.TODO(),filter, bson.M{"$set": docs})

	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"items":   result,
	})
}