package data

import (
	"context"
	"errors"
	"sync"
	"task_manager3/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
	mutex      sync.Mutex
}

func NewTaskService(db *mongo.Database) *TaskService {
	return &TaskService{
		collection: db.Collection("tasks"),
		mutex:      sync.Mutex{},
	}
}

func (tc *TaskService) CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	if task.Title == "" {
		return nil, errors.New("title can not be empty")
	}

	if task.Status == "" {
		return nil, errors.New("status can not be empty")
	}

	if task.Status != "Complete" && task.Status != "Not Started" && task.Status != "In Progress" {
		return nil, errors.New("status must be Compelete, Not Started or In Progress")
	}
	return tc.collection.InsertOne(context.TODO(), task)
}

func (tc *TaskService) GetTasks() ([]models.Task, error) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	cursor, err := tc.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	if err = cursor.All(context.TODO(), &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tc *TaskService) GetTask(id string) (*models.Task, error) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = tc.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (tc *TaskService) UpdateTask(id string, update models.Task) (*mongo.UpdateResult, error) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return tc.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.D{
			{Key: "$set", Value: update},
		},
	)
}

func (tc *TaskService) DeleteTask(id string) (*mongo.DeleteResult, error) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return tc.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
}
