package users

import (
	"backend/db/migration/users"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS = "users"

func GetUsers(c *fiber.Ctx, db *mongo.Database) error {
	var filter = bson.D{{}}
	
	if c.Params("id") != "" {
		objectId, err := primitive.ObjectIDFromHex(c.Params("id"))
		
		if err != nil {
			return err
		}
		filter = bson.D{{Key: "_id", Value: objectId}}
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
		"items":   results,
	})
}

func AddUser(c *fiber.Ctx, db *mongo.Database) error {
	userDb := db.Collection(USERS)

	data := new(users.Users)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	docs := []interface{}{
		&users.Users{
			Id:				primitive.NewObjectID(),
			Email:		data.Email,
			Name:			data.Name,
			Password:	data.Password,
		}}

	result, err := userDb.InsertMany(context.TODO(), docs)

	if err != nil {
		panic(err)
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"items":   result,
	})

}