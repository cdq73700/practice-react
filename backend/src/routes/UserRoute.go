package routes

import (
	api_v1 "backend/src/api/v1/user"
	"backend/src/handler"
	"backend/src/models"

	"github.com/gofiber/fiber/v2"
)

// Represents our handler with our use-case / service.
type UserHandler struct {
	userService models.UserService
}

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
	userHandler := api_v1.NewUserHandler(userRepository)

	// Create a handler based on our created service / use-case.
	handle := &UserHandler{
		userService: &userHandler,
	}

	userRoute.Get("", handle.userService.GetUserList)
	userRoute.Get("/:Id", handle.userService.GetUser)
	userRoute.Post("/create", handle.userService.CreateUser)
	userRoute.Put("/update", handle.userService.UpdateUser)
	userRoute.Post("/delete", handle.userService.DeleteUser)
}