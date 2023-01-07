package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserList(c *fiber.Ctx, db *mongo.Database) error {
	list := db.Collection(USERS)

	cursor, err := list.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}

	return Response(c, cursor, err)
}