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
	GetUser(*fiber.Ctx, primitive.ObjectID) error
	CreateUser(*fiber.Ctx) error
	UpdateUser(*fiber.Ctx, *UserStruct) error
	DeleteUser(*fiber.Ctx, primitive.ObjectID) error
}

type UserService interface {
	GetUserList(*fiber.Ctx) error
	GetUser(*fiber.Ctx) error
	CreateUser(*fiber.Ctx) error
	UpdateUser(*fiber.Ctx) error
	DeleteUser(*fiber.Ctx) error
}
