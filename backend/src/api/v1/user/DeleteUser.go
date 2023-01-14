package api_v1

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *userService) DeleteUser(c *fiber.Ctx) error {

	user, err := service.checkIfUserExistsMiddleware(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return service.repository.DeleteUser(c, user.Id)
}

func (db *dbStruct) DeleteUser(c *fiber.Ctx, Id primitive.ObjectID) error {

	// filter
	filter := bson.D{{Key: "_id", Value: Id}}

	result, err := db.model.Collection.DeleteOne(context.TODO(),filter)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"items": Id,
		"result": result,
	})
}