package users

import (
	"backend/db/migration/users"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS = "users"

func GetUsers(c *fiber.Ctx, db *mongo.Database) error {
	var filter = bson.D{{}}
	if c.Params("name") != "" {
		filter = bson.D{{Key: "username", Value: c.Params("name")}}
	}

	userList := db.Collection(USERS)
	cursor, err := userList.Find(context.TODO(), filter)
	if err != nil {
		return err
	}

	var results []users.Users
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	if len(results) == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   "There are no users!",
		})
	}
	return c.JSON(&fiber.Map{
		"success": true,
		"users":   results,
	})
}
