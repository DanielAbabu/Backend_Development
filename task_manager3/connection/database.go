package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	"task_manager3/controllers"
	"task_manager3/data"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateConnection() (*controllers.TaskController, *controllers.UserController) {
	// Retrieve the MongoDB URI from the environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database("taskmanager3")
	taskService := data.NewTaskService(db)
	taskController := controllers.NewTaskController(taskService)
	userService := data.NewUserService(db)
	userController := controllers.NewUserController(userService)

	return taskController, userController
}
