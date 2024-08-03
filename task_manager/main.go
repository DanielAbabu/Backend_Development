package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
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

	db := client.Database("taskmanager")
	taskService := data.NewTaskService(db)
	taskController := controllers.NewTaskController(taskService)

	r := router.SetupRouter(taskController)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
