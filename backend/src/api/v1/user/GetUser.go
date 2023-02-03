package api_v1

import (
	"backend/src/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *userService) GetUser(c *fiber.Ctx) error {

	targetedUserID, err := primitive.ObjectIDFromHex(c.Params("Id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID!",
		})
	}
	
	return service.repository.GetUser(c,targetedUserID)
}

func (db *dbStruct) GetUser(c *fiber.Ctx,id primitive.ObjectID) error {
	user := &models.UserStruct{}

	// filter := bson.D{{Key: "_id", Value: targetedUserID}}
	filter := bson.D{{Key: "_id", Value: id}}

	res := db.model.Collection.FindOne(context.TODO(), filter)

	if err := res.Decode(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"items":   &user,
	})
}