package api_v1

import (
	"backend/src/models"

	"github.com/gofiber/fiber/v2"
)

type dbStruct struct {
	model *models.MongoStrict
}
type userService struct {
	repository models.UserRepository
}
type userRepository models.UserRepository
type userStruct models.UserStruct



// Create a new repository with MongoDB as the driver.
func NewUserRepository(mongoDbConnection *models.MongoStrict) userRepository {
	return &dbStruct{model: mongoDbConnection}
}

// Create a new handler.
func NewUserHandler(repository userRepository) userService {
	return userService{repository: repository}
}

func (service *userService)checkIfUserExistsMiddleware(c *fiber.Ctx) (*models.UserStruct, error) {
	var err error
	data := new(models.UserStruct)

	// Fetch parameter.
	err = c.BodyParser(data)
	if err != nil {
		return nil, c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Please specify a valid user ID!",
		})
	}

	// Check if user exists.
	err = service.repository.GetUser(c, data.Id)
	if err != nil {
		return nil, c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	return data, nil
}