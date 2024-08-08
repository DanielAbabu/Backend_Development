package data

import (
	"context"
	"errors"
	"task_manager3/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	collection := db.Collection("tasks")
	return &TaskService{collection: collection}

}

func (tc *TaskService) CreateTask(newtask *models.Task) (*mongo.InsertOneResult, error) {

	if newtask.Title == "" {
		return nil, errors.New("title can not be empty")
	}

	if newtask.Status == "" {
		return nil, errors.New("status can not be empty")
	}
	nn, err := tc.collection.InsertOne(context.TODO(), newtask)
	if err != nil {
		return nil, errors.New("couldnt insert")
	}

	return nn, nil
}

func (tc *TaskService) GetTask(id string) (*models.Task, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task

	err = tc.collection.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&task)

	if err != nil {
		return nil, err
	}

	return &task, nil

}
func (tc *TaskService) GetTasks(userid string) (*[]models.Task, error) {
	uid, err := primitive.ObjectIDFromHex(userid)
	if err != nil {
		return nil, err
	}
	cursor, err := tc.collection.Find(context.TODO(), bson.M{"user_id": uid})

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task

	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}
	return &tasks, nil

}
func (tc *TaskService) UpdateTask(id string, updatedtask *models.Task) (*models.Task, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	temTask, _ := tc.GetTask(id)
	updatedtask.ID = temTask.ID
	updatedtask.UserID = temTask.UserID

	_, err = tc.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": oid},
		bson.D{
			{Key: "$set", Value: updatedtask},
		},
	)

	if err != nil {
		return nil, err
	}

	return updatedtask, nil
}
func (tc *TaskService) RemoveTask(id string) (*mongo.DeleteResult, error) {

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return tc.collection.DeleteOne(context.TODO(), bson.M{"_id": oid})

}
