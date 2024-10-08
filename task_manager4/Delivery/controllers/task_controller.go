package controllers

import (
	"net/http"
	"task_manager4/Domain"

	"github.com/gin-gonic/gin"
)

func GetAllTasks(tu Domain.TaskUsecase, c *gin.Context) {
	userid, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "user id doesn't exsist"})
		return
	}
	usr, ok := userid.(Domain.DBUser)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "type assertion didn't work"})
		return
	}
	tasks, err := tu.GetTasks(usr.ID.Hex())
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, tasks)

}

func GetTaskById(tu Domain.TaskUsecase, c *gin.Context) {
	id := c.Param("id")
	userid, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "user id doesn't exsist"})
		return
	}
	usr, ok := userid.(Domain.DBUser)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "type assertion didn't work"})
		return
	}

	task, err := tu.GetTask(id, usr.ID.Hex())
	if err == nil {
		c.IndentedJSON(http.StatusOK, task)
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
}

func PostTask(tu Domain.TaskUsecase, c *gin.Context) {
	var task Domain.Task

	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})

	}
	user, exists := c.Get("user")
	if !exists {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": "user couldn't be found"})

	}

	usr, ok := user.(Domain.DBUser)
	if !ok {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": "type assertion didn't work"})

	}

	task, err := tu.PostTask(task, usr)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})

	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "created sucessfully", "task": task})

}

func DeleteTask(tu Domain.TaskUsecase, c *gin.Context) {
	id := c.Param("id")
	userid, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "user id doesn't exsist"})
		return
	}
	usr, ok := userid.(Domain.DBUser)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "type assertion didn't work"})
		return
	}

	err := tu.DeleteTask(id, usr.ID.Hex())
	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"messages": "deleted successfully"})

	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})

}

func UpdateTask(tu Domain.TaskUsecase, c *gin.Context) {
	var task Domain.Task

	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}

	id := c.Param("id")
	user, exists := c.Get("user")
	if !exists {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": "user couldn't be found"})

	}

	usr, ok := user.(Domain.DBUser)
	if !ok {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": "type assertion didn't work"})

	}

	task, err := tu.UpdateTask(id, task, usr)
	if err == nil {
		c.IndentedJSON(http.StatusOK, task)

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "task not found"})

}
