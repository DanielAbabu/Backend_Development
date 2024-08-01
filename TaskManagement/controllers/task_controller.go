package controllers

import (
	"TaskManagement/data"
	"TaskManagement/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTask(g *gin.Context) {
	id := g.Param("id")

	task, err := data.Get(id)

	if err != nil {
		g.JSON(http.StatusNotFound, error.Error(err))
		return
	}

	g.JSON(http.StatusFound, task)
}

func GetAllTask(g *gin.Context) {
	task := data.GetAll()

	if len(task) == 0 {
		g.JSON(http.StatusFound, "empty task")
		return
	}
	g.JSON(http.StatusFound, task)
}

func AddTask(g *gin.Context) {
	var task models.Task

	if err := g.BindJSON(&task); err != nil {
		g.JSON(http.StatusBadRequest, error.Error(err))
		return
	}

	ntask := data.Add(task)
	g.JSON(http.StatusAccepted, gin.H{"The New task": ntask})
}

func UpdateTask(g *gin.Context) {
	var task models.Task
	id := g.Param("id")

	if err := g.BindJSON(&task); err != nil {
		g.JSON(http.StatusBadRequest, error.Error(err))
		return
	}

	if err := data.Update(id, task); err != nil {
		g.JSON(http.StatusNotFound, error.Error(err))
		return
	}

	g.JSON(http.StatusAccepted, "Task Updated")
}

func DeleteTask(g *gin.Context) {
	id := g.Param("id")

	err := data.Delete(id)

	if err != nil {
		g.JSON(http.StatusNotFound, error.Error(err))
		return
	}

	g.JSON(http.StatusFound, "task Deleted")
}
