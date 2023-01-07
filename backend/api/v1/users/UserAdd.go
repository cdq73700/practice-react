package users

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UserAdd(c *fiber.Ctx, db *mongo.Database) error {
	userDb := db.Collection(USERS)

	data := new(UserStruct)

	if err := c.BodyParser(data); err != nil {
		return err
	}

	docs := []interface{}{
		&UserStruct{
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