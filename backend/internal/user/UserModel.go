package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const USERS = "users"

var UserCollectionClient *mongo.Collection

type UserStruct struct {
	Id				primitive.ObjectID	"_id"
	Email			string
	Name			string
	Password	[]byte
} 

// Our repository will implement these methods.
type UserRepository interface {
	GetUserList(ctx context.Context) (*[]UserStruct, error)
	GetUser(ctx context.Context, Id primitive.ObjectID) (*UserStruct, error)
	CreateUser(ctx context.Context, user *UserStruct) error
	UpdateUser(ctx context.Context, Id primitive.ObjectID, user *UserStruct) error
	DeleteUser(ctx context.Context, Id primitive.ObjectID) error
}

type UserService interface {
	GetUserList(ctx context.Context) (*[]UserStruct, error)
	GetUser(ctx context.Context, Id primitive.ObjectID) (*UserStruct, error)
	CreateUser(ctx context.Context, user *UserStruct) error
	UpdateUser(ctx context.Context, Id primitive.ObjectID, user *UserStruct) error
	DeleteUser(ctx context.Context, Id primitive.ObjectID) error
}

func CreatePassword(password string) []byte {
	hash,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	return hash
}
