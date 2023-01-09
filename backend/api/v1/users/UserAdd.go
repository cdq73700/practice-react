package users

import (
	model "backend/model"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UserAdd(c *fiber.Ctx) error {
	userDb := model.UserCollectionClient

	data := new(model.UserStruct)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	docs := []interface{}{
		&model.UserStruct{
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