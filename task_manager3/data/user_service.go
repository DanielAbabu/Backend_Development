package data

import (
	"context"
	"errors"

	"task_manager3/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		collection: db.Collection("user"),
	}
}

func (us *UserService) Register(user *models.User) error {
	var usr models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&usr)

	if err == nil {
		return errors.New("Account Already exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	_, err = us.collection.InsertOne(context.TODO(), user)

	return err
}

func (us *UserService) Authenticate(email string, password string) (*models.User, error) {
	var user models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return &user, nil

}
