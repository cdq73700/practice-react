package users

import (
	model "backend/Model"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UserDelete(c *fiber.Ctx) error {
	data := new(model.UserStruct)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: data.Id}}

	items := model.UserCollectionClient

	result, err := items.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"items":   result,
	})
}