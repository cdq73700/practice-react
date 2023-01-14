package api_v1

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (db *dbStruct) CreateUser(c *fiber.Ctx) error {

	var err error
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

	_, err = db.model.Collection.InsertOne(context.TODO(), &user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": &user,
	})
}

func createPassword(password string) ([]byte, error) {
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}