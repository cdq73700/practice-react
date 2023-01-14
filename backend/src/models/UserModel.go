package models

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStruct struct {
	Id primitive.ObjectID `bson:"_id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	Password []byte `bson:"password"`
}

type UserRepository interface {
	GetUserList(*fiber.Ctx) error
	GetUser(*fiber.Ctx) error
	CreateUser(*fiber.Ctx) error
}
