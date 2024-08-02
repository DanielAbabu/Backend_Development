package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.TaskController) *gin.Engine {
	r := gin.Default()

	// Task routes
	r.GET("/tasks", tc.GetTasks)
	r.GET("/tasks/:id", tc.GetTask)
	r.POST("/tasks", tc.CreateTask)
	r.PUT("/tasks/:id", tc.UpdateTask)
	r.DELETE("/tasks/:id", tc.DeleteTask)

	return r
}
