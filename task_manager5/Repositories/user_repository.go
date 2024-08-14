package Repositories

import (
	"context"
	"errors"
	"task_manager5/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	coll *mongo.Collection
}

func (UR *UserRepo) EnsureIndexes() error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := UR.coll.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func NewUserRepo(db *mongo.Database, name string) (*UserRepo, error) {
	UR := &UserRepo{
		coll: db.Collection(name),
	}

	// Ensure indexes are created
	if err := UR.EnsureIndexes(); err != nil {
		return nil, err
	}

	return UR, nil
}

func (UR *UserRepo) FindByEmail(email string) (Domain.UserInput, error) {
	var userDB Domain.UserInput
	query := bson.M{"email": email}

	err := UR.coll.FindOne(context.TODO(), query).Decode(&userDB)
	if err != nil {
		return Domain.UserInput{}, err
	}

	return userDB, nil
}

func (UR *UserRepo) FindById(id string) (Domain.UserInput, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}
	var user Domain.UserInput
	err := UR.coll.FindOne(context.TODO(), query).Decode(&user)
	if err != nil {
		return Domain.UserInput{}, err
	}

	return user, nil
}

func (UR *UserRepo) FindAllUsers() ([]Domain.DBUser, error) {
	cursor, err := UR.coll.Find(context.TODO(), bson.M{})

	if err != nil {
		return nil, err
	}

	var users []Domain.DBUser

	for cursor.Next(context.TODO()) {
		user := Domain.UserInput{}
		err := cursor.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, Domain.ChangeToOutput(user))
	}

	return users, nil
}

func (UR *UserRepo) UpdateUserById(id string, user Domain.UserInput, is_admin bool) (Domain.DBUser, error) {
	obId, _ := primitive.ObjectIDFromHex(id)
	user.ID = obId
	user.IsAdmin = is_admin
	bsonModel, err := bson.Marshal(user)
	if err != nil {
		return Domain.DBUser{}, err
	}

	var doc bson.M
	err = bson.Unmarshal(bsonModel, &doc)
	if err != nil {
		return Domain.DBUser{}, err
	}
	filter := bson.D{{Key: "_id", Value: obId}, {Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: doc}}

	_, err = UR.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return Domain.DBUser{}, err
	}

	return Domain.ChangeToOutput(user), nil
}

func (UR *UserRepo) CreateUser(user Domain.UserInput) (Domain.DBUser, error) {
	// Set other user properties
	user.ID = primitive.NewObjectID()
	user.IsAdmin = false

	// Insert the new user
	_, er := UR.coll.InsertOne(context.TODO(), &user)
	if er != nil {
		// Check if the error is due to a duplicate key
		if mongo.IsDuplicateKeyError(er) {
			return Domain.DBUser{}, errors.New("email already exists")
		}
		return Domain.DBUser{}, er
	}

	return Domain.ChangeToOutput(user), nil

}

func (UR *UserRepo) DeleteUserByID(id string) error {
	obId, _ := primitive.ObjectIDFromHex(id)
	query := bson.M{"_id": obId}

	res, err := UR.coll.DeleteOne(context.TODO(), query)

	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("no document with this id exists")
	}

	return nil
}
