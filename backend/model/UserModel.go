package model

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const USERS = "users"

var UserCollectionClient *mongo.Collection

type UserStruct struct {
	Id				primitive.ObjectID	"_id"
	Email			string
	Name			string
	Password	[]byte
} 

func UserInitialized() {
	UserCollectionClient = GetMongoDbCollection(USERS)

}

func CreatePassword(password string) []byte {
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	return hash
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