package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Represents that we will use mongoDB in order to implement the methods.
type mongoDBRepository struct {
	mongodb *mongo.Collection
}

// Create a new repository with MongoDB as the driver.
func NewUserRepository(mongoDbConnection *mongo.Collection) UserRepository {
	return &mongoDBRepository{
		mongodb: mongoDbConnection,
	}
}

// CreateUser implements UserRepository
func (r *mongoDBRepository) CreateUser(ctx context.Context, user *UserStruct) error {
	_, err := r.mongodb.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}

// DeleteUser implements UserRepository
func (*mongoDBRepository) DeleteUser(ctx context.Context, Id primitive.ObjectID) error {
	panic("unimplemented")
}

// GetUser implements UserRepository
func (r *mongoDBRepository) GetUser(ctx context.Context, Id primitive.ObjectID) (*UserStruct, error) {
	// Initialize variables.
	user := &UserStruct{}

	// filter
	filter := bson.D{{Key: "_id", Value: Id}}

	// Get user.
	res := r.mongodb.FindOne(ctx, filter)

	if err := res.Decode(&user); err != nil {
		return nil, err
	}

	// Return user.
	return user, nil
}

// GetUserList implements UserRepository
func (r *mongoDBRepository) GetUserList(ctx context.Context) (*[]UserStruct, error) {
	// Initialize variables.
	var users []UserStruct

	res, err := r.mongodb.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer res.Close(context.TODO())

	if err = res.All(context.TODO(), &users); err != nil {
		return nil, err
	}

	// Return all of our users.
	return &users, nil
}

// UpdateUser implements UserRepository
func (*mongoDBRepository) UpdateUser(ctx context.Context, Id primitive.ObjectID, user *UserStruct) error {
	panic("unimplemented")
}
