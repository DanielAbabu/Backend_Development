package main

import (
	"task_manager3/connection"
	"task_manager3/router"
)

func main() {
	taskcontroller, usercontroller := connection.CreateConnection()

	router := router.SetRouter(taskcontroller, usercontroller)
	router.Run(":8080")
}
