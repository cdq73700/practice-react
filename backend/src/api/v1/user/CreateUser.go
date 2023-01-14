package api_v1

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (service *userService) CreateUser(c *fiber.Ctx) error {
	return service.repository.CreateUser(c)
}

func (db *dbStruct) CreateUser(c *fiber.Ctx) error {
	var err error
	var result *mongo.InsertOneResult
	user := &userStruct{}

	err = c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	password, err := createPassword(string(user.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	
	user.Id = primitive.NewObjectID()
	user.Password = password

	result, err = db.model.Collection.InsertOne(context.TODO(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": &user,
		"result": &result,
	})
}

func createPassword(password string) ([]byte, error) {
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}