package controllers

import (
	"net/http"
	"task_manager3/data"
	"task_manager3/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	service *data.TaskService
}

func NewTaskController(service *data.TaskService) *TaskController {
	return &TaskController{service: service}
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var newtask models.Task

	if err := c.BindJSON(&newtask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload", "details": err.Error()})
		return
	}
	// log.Printf("Received Task: %+v\n", newtask)

	userid, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id not found"})
		return
	}
	useridstr, ok := userid.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID is not valid"})
		return
	}
	userObjectID, err := primitive.ObjectIDFromHex(useridstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "userID is not a valid ObjectID"})
		return
	}
	newtask.UserID = userObjectID

	result, err := tc.service.CreateTask(&newtask)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		newtask.ID = oid
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive ID"})
		return
	}
	c.JSON(http.StatusCreated, newtask)

}
func (tc *TaskController) GetTask(c *gin.Context) {

	id := c.Param("id")
	task, err := tc.service.GetTask(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)

}
func (tc *TaskController) GetTasks(c *gin.Context) {

	userid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "userID doesn't exsist"})
		return
	}

	userID := userid.(string)
	tasks, err := tc.service.GetTasks(userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})

}
func (tc *TaskController) UpdateTask(c *gin.Context) {

	id := c.Param("id")

	var updatedTask models.Task

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := tc.service.UpdateTask(id, &updatedTask)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Updated task": updated})

}
func (tc *TaskController) RemoveTask(c *gin.Context) {

	id := c.Param("id")
	result, err := tc.service.RemoveTask(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": result})

}
