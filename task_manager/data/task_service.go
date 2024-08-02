package data

import (
	"context"
	"log"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://danielababu:q6NBfCGOlAOAuR4F@taskmanagement.ntpfnxc.mongodb.net/task_management?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("task_management").Collection("tasks")
}

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	return &TaskService{
		collection: db.Collection("tasks"),
	}
}

// CreateTask inserts a new task into the database
func (tc *TaskService) CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context.TODO(), task)
}

// GetTasks retrieves all tasks from the database
func (tc *TaskService) GetTasks() ([]models.Task, error) {
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
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

// GetTask retrieves a task by ID from the database
func (tc *TaskService) GetTask(id string) (*models.Task, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

// UpdateTask updates a task by ID in the database
func (tc *TaskService) UpdateTask(id string, update models.Task) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.D{
			{Key: "$set", Value: update},
		},
	)
}

// DeleteTask deletes a task by ID from the database
func (tc *TaskService) DeleteTask(id string) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
}
