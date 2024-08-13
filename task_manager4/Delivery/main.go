package main

import (
	"context"
	"fmt"
	"log"
	"task_manager4/Delivery/connection"
	"task_manager4/Delivery/routers"
)

func main() {
	fmt.Println("Server started")
	client := connection.CreateConnection()

	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	db := client.Database("taskmanager4")
	routers.Run(db)

}
