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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Connected to MongoDB!")
	db := client.Database("taskmanager3")
	userService := data.NewUserService(db)
	taskService := data.NewTaskService(db)
	taskController := controllers.NewTaskController(taskService)
	userController := controllers.NewUserController(*userService)

	return taskController, userController
}
