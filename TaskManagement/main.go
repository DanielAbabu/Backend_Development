package main

import (
	"TaskManagement/router"
)

func main() {
	route := router.Setup()

	route.Run("localhost:8080")
}
