package api_v1

import (
	"backend/src/models"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *userService) UpdateUser(c *fiber.Ctx) error {

	user, err := service.checkIfUserExistsMiddleware(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return service.repository.UpdateUser(c, user)
}

func (db *dbStruct) UpdateUser(c *fiber.Ctx, user *models.UserStruct) error {

	// filter
	filter := bson.D{{Key: "_id", Value: user.Id}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: user.Name},{Key: "email", Value: user.Email}}}}


	result, err := db.model.Collection.UpdateOne(context.TODO(),filter, update)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"items": user,
		"update": update,
		"result": result,
	})
}