package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserDelete(c *fiber.Ctx, db *mongo.Database) error {
	data := new(UserStruct)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	filter := bson.D{{Key: "_id", Value: data.Id}}

	items := db.Collection(USERS)

	result, err := items.DeleteOne(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"items":   result,
	})
}