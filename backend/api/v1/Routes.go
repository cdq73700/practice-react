package v1

import (
	v1 "backend/api/v1/user"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoDBRepository struct{
	*mongo.Collection
}

// Creates a new handler.
func NewUserHandler(userRoute fiber.Router,collection *mongo.Collection) {

	repository := &mongoDBRepository{Collection: collection}

	// Declare routing endpoints for general routes.
	userRoute.Get("", repository.getUserList)

	// Declare routing endpoints for specific routes.
	// userRoute.Get("/:userID", handle.getUser)
	// userRoute.Post("/create", handle.createUser)
	// userRoute.Put("/update", handle.checkIfUserExistsMiddleware, handle.updateUser)
	// userRoute.Delete("/delete", handle.checkIfUserExistsMiddleware, handle.deleteUser)
}

// Gets all users.
func (repository *mongoDBRepository)getUserList(c *fiber.Ctx) error {

	// Get all users.
	users, err := v1.GetUserList(repository.Collection)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"items":   users,
	})
}