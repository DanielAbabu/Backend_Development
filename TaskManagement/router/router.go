package router

import (
	"TaskManagement/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTask)
	router.GET("/tasks/:id", controllers.GetTask)
	router.POST("/tasks", controllers.AddTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)

	return router

}
