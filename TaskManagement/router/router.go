package router

import (
	"TaskManagement/controllers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTask)
	router.GET("/task/:id", controllers.GetTask)
	router.POST("/task", controllers.AddTask)
	router.DELETE("/task/:id", controllers.DeleteTask)
	router.PUT("/task/:id", controllers.UpdateTask)

	return router

}
