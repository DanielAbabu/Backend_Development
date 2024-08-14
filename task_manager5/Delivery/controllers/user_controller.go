package controllers

import (
	"net/http"
	"task_manager5/Domain"

	"github.com/gin-gonic/gin"
)

func Register(uu Domain.UserUsecase, c *gin.Context) {
	var user Domain.UserInput

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usr, er := uu.Signup(user)
	if er != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": er.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User created Successfully.", "user": usr})
}

func Login(uu Domain.UserUsecase, c *gin.Context) {
	var user Domain.UserInput
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	usr, token, err := uu.Login(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "user": usr})
}

func GetAllUsers(uu Domain.UserUsecase, c *gin.Context) {
	users, err := uu.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, users)
}

func GetUserById(uu Domain.UserUsecase, c *gin.Context) {
	id := c.Param("id")
	user, err := uu.GetUser(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, user)
}

func MakeAdmin(uu Domain.UserUsecase, c *gin.Context) {
	id := c.Param("id")
	user, err := uu.MakeAdmin(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, user)
	return
}

func DeleteUser(uu Domain.UserUsecase, c *gin.Context) {
	id := c.Param("id")
	err := uu.DeleteUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "User deleted successfully"})
}

func UpdateUser(uu Domain.UserUsecase, c *gin.Context) {
	id := c.Param("id")
	var user Domain.UserInput
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr, er := uu.UpdateUser(id, user)
	if er != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"error": er.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, usr)
}
