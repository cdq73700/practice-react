package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Implementation of the repository in this service.
type userService struct {
	userRepository UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

// Implementation of 'GetUserList'.
func (s *userService) GetUserList(ctx context.Context) (*[]UserStruct, error) {
	return s.userRepository.GetUserList(ctx)
}

// Implementation of 'GetUser'.
func (s *userService) GetUser(ctx context.Context, userID primitive.ObjectID) (*UserStruct, error) {
	return s.userRepository.GetUser(ctx, userID)
}

// Implementation of 'CreateUser'.
func (s *userService) CreateUser(ctx context.Context, user *UserStruct) error {
	// Set default value of 'Created' and 'Modified'.
	user.Id = primitive.NewObjectID()
	user.Password = CreatePassword(string(user.Password))
	//user.Created = time.Now().Unix()
	//user.Modified = time.Now().Unix()

	// Pass to the repository layer.
	return s.userRepository.CreateUser(ctx, user)
}

// Implementation of 'UpdateUser'.
func (s *userService) UpdateUser(ctx context.Context, userID primitive.ObjectID, user *UserStruct) error {
	// Set value for 'Modified' attribute.
	//user.Modified = time.Now().Unix()

	// Pass to the repository layer.
	return s.userRepository.UpdateUser(ctx, userID, user)
}

// Implementation of 'DeleteUser'.
func (s *userService) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	return s.userRepository.DeleteUser(ctx, userID)
}