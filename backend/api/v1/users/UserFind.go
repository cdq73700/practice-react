package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserFind(c *fiber.Ctx, db *mongo.Database) error {
	var filter = bson.D{{}}
	
	if c.Params("id") != "" {
		objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
		
		if err != nil {
			return err
		}
		filter = bson.D{{Key: "_id", Value: objectId}}
	}

	items := db.Collection(USERS)

	cursor, err := items.Find(context.TODO(), filter)
	if err != nil {
		return err
	}

	return Response(c, cursor, err)
}