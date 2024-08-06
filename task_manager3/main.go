package main

import (
	"log"
	"task_manager3/connection"
	"task_manager3/router"
)

func main() {

	r := router.SetupRouter(connection.CreateConnection())
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
