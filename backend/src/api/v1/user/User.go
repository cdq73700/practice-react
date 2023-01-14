package api_v1

import (
	"backend/src/models"
)

type dbStruct struct {
	model *models.MongoStrict}
type userStruct models.UserStruct
type userRepository models.UserRepository

// Create a new repository with MongoDB as the driver.
func NewUserRepository(mongoDbConnection *models.MongoStrict) userRepository {
	return &dbStruct{model: mongoDbConnection}
}
