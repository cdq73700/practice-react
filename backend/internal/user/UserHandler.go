package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Represents our handler with our use-case / service.
type UserHandler struct {
	userService UserService
}

// Creates a new handler.
func NewUserHandler(userRoute fiber.Router, us UserService) {
	// Create a handler based on our created service / use-case.
	handle := &UserHandler{
		userService: us,
	}

	// Declare routing endpoints for general routes.
	userRoute.Get("", handle.getUserList)

	// Declare routing endpoints for specific routes.
	userRoute.Get("/:userID", handle.getUser)
	userRoute.Post("/add", handle.createUser)
	userRoute.Put("/update", handle.checkIfUserExistsMiddleware, handle.updateUser)
	userRoute.Delete("/delete", handle.checkIfUserExistsMiddleware, handle.deleteUser)
}

// Gets all users.
func (h *UserHandler) getUserList(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Get all users.
	users, err := h.userService.GetUserList(customContext)
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

// Gets a single user.
func (h *UserHandler) getUser(c *fiber.Ctx) error {
	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Fetch parameter.
	targetedUserID, err := primitive.ObjectIDFromHex(c.Params("userID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID!",
		})
	}

	// Get one user.
	user, err := h.userService.GetUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return results.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"items":   user,
	})
}

// Creates a single user.
func (h *UserHandler) createUser(c *fiber.Ctx) error {
	// Initialize variables.
	user := &UserStruct{}

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Create one user.
	err := h.userService.CreateUser(customContext, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"status":  "success",
		"message": "User has been created successfully!",
	})
}

// Updates a single user.
func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	// Initialize variables.
	user := &UserStruct{}
	targetedUserID,err := primitive.ObjectIDFromHex(c.Locals("_id").(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse request body.
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Update one user.
	err = h.userService.UpdateUser(customContext, targetedUserID, user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return result.
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "User has been updated successfully!",
	})
}

// Deletes a single user.
func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	// Initialize previous user ID.
	targetedUserID,err := primitive.ObjectIDFromHex(c.Locals("_id").(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Create cancellable context.
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Delete one user.
	err = h.userService.DeleteUser(customContext, targetedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	// Return 204 No Content.
	return c.SendStatus(fiber.StatusNoContent)
}