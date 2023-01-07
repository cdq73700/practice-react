package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const USERS = "users"

type UserStruct struct {
	Id				primitive.ObjectID	"_id"
	Email			string
	Name			string
	Password	string
} 

func (user *UserStruct) Initialized(Id primitive.ObjectID,Email string, Name string, Password string) {
	id := primitive.NewObjectID()
	if Id != primitive.NilObjectID {
		id = Id
	}
	user.Id = id
	user.Email = Email
	user.Name = Name
	user.Password = Password
}

func Response(c *fiber.Ctx,cursor *mongo.Cursor,err error) error {
	var results []UserStruct
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
