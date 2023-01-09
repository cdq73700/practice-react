package users

import (
	model "backend/Model"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserFind(c *fiber.Ctx) error {
	var filter = bson.D{{}}
	
	if c.Params("id") != "" {
		objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
		
		if err != nil {
			return err
		}
		filter = bson.D{{Key: "_id", Value: objectId}}
	}

	items := model.UserCollectionClient

	cursor, err := items.Find(context.TODO(), filter)
	if err != nil {
		return err
	}

	return model.Response(c, cursor, err)
}