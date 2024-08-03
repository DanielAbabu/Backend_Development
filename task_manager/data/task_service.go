package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// var collection *mongo.Collection

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Retrieve the MongoDB URI from the environment variable
// 	mongoURI := os.Getenv("MONGODB_URI")

// 	clientOptions := options.Client().ApplyURI(mongoURI)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection = client.Database("task_management").Collection("tasks")
// }

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(db *mongo.Database) *TaskService {
	return &TaskService{
		collection: db.Collection("tasks"),
	}
}

func (tc *TaskService) CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
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
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	return tc.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
}
