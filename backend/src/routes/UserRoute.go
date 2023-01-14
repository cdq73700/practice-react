package routes

import (
	api_v1 "backend/src/api/v1/user"
	"backend/src/handler"

	"github.com/gofiber/fiber/v2"
)

func NewUserRoutes(userRoute fiber.Router) {

	db, err := handler.NewDatabase("test","users")

	if err != nil {
		userRoute.Get("", func(ctx * fiber.Ctx) error {
			return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			}) 
		 })
	}

	userRepository := api_v1.NewUserRepository(db)

	userRoute.Get("", userRepository.GetUserList)
	userRoute.Get("/:Id", userRepository.GetUser)
	userRoute.Post("/create", userRepository.CreateUser)
}