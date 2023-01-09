package users

import (
	model "backend/Model"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UserUpdate(c *fiber.Ctx) error {
	userDb := model.UserCollectionClient

	data := new(model.UserStruct)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: data.Id}}
	
	docs := &model.UserStruct {
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