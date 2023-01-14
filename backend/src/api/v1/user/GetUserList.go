package api_v1

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *userService) GetUserList(c *fiber.Ctx) error {
	return service.repository.GetUserList(c)
}

func (db *dbStruct) GetUserList(c *fiber.Ctx) error {
	var users []userStruct

	res, err := db.model.Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		}) 
	}
	defer res.Close(context.TODO())

	if err = res.All(context.TODO(), &users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		}) 
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"items":   &users,
	})
}