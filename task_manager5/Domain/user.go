package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserInput struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password" bson:"password"`
	IsAdmin  bool               `json:"isadmin" bson:"isadmin"`
}

type DBUser struct {
	ID      primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name,omitempty"`
	Email   string             `json:"email,omitempty" bson:"email,omitempty"`
	IsAdmin bool               `json:"isadmin" bson:"isadmin"`
}

type UserUsecase interface {
	Login(UserInput) (DBUser, string, error)
	Signup(UserInput) (DBUser, error)
	GetUsers() ([]DBUser, error)
	GetUser(string) (DBUser, error)
	MakeAdmin(string) (DBUser, error)
	UpdateUser(string, UserInput) (DBUser, error)
	DeleteUser(string) error
}

type UserRepository interface {
	FindByEmail(string) (UserInput, error)
	FindById(string) (UserInput, error)
	FindAllUsers() ([]DBUser, error)
	UpdateUserById(string, UserInput, bool) (DBUser, error)
	CreateUser(UserInput) (DBUser, error)
	DeleteUserByID(string) error
}

func ChangeToOutput(user UserInput) DBUser {
	var Ouser DBUser
	Ouser.ID = user.ID
	Ouser.Name = user.Name
	Ouser.Email = user.Email
	Ouser.IsAdmin = user.IsAdmin

	return Ouser
}
