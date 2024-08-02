package main

import (
	"context"
	"fmt"
	"log"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"

	// "github.com/danielababu/task_manager/controllers"
	// "github.com/danielababu/task_manager/data"
	// "github.com/danielababu/task_manager/router"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://danielababu:q6NBfCGOlAOAuR4F@taskmanagement.ntpfnxc.mongodb.net/?retryWrites=true&w=majority")

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

	// Get a handle for your database
	db := client.Database("taskmanager")

	// Initialize task service and controller
	taskService := data.NewTaskService(db)
	taskController := controllers.NewTaskController(taskService)

	// Set up the router
	r := router.SetupRouter(taskController)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
