package users

import (
	model "backend/model"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func UserList(c *fiber.Ctx) error {
	list := model.UserCollectionClient

	cursor, err := list.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return err
	}

	return model.Response(c, cursor, err)
}