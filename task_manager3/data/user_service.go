package data

import (
	"context"
	"errors"
	"log"
	"os"
	"task_manager3/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *UserService {
	collection := db.Collection("users")
	return &UserService{collection: collection}

}

func (us *UserService) Register(user *models.User) (*mongo.InsertOneResult, error) {
	user.ID = primitive.NewObjectID()
	if user.Email == "" {
		return nil, errors.New("email can not be empty")
	}
	var u models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&u)

	if err == nil {
		return nil, errors.New("email already exists")
	}
	if user.Role == "" {
		return nil, errors.New("role can not be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)
	return us.collection.InsertOne(context.TODO(), user)
}

func (us *UserService) Login(user *models.User) (string, error) {

	var userfound models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&userfound)

	if err != nil {
		return "", err
	}
	log.Printf("Hashed password: %s\n", userfound.Password)
	log.Printf("Input password: %s\n", user.Password)

	err = bcrypt.CompareHashAndPassword([]byte(userfound.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userfound.ID.Hex(),
		"email":   userfound.Email,
		"role":    userfound.Role,
	})

	var jwtSecret []byte = []byte(os.Getenv("secret"))

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return jwtToken, nil

}

func (us *UserService) GetUser(email string) (*models.User, error) {
	var user models.User
	err := us.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (us *UserService) GetUsers() (*[]models.User, error) {
	cursor, err := us.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	var users []models.User

	if err = cursor.All(context.TODO(), &users); err != nil {
		return nil, err
	}
	return &users, nil
}

func (us *UserService) DeleteUser(email string) error {
	_, err := us.collection.DeleteOne(context.TODO(), bson.M{"email": email})
	return err
}
