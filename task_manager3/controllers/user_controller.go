package controllers

import (
	"log"
	"net/http"
	"task_manager3/data"
	"task_manager3/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	service data.UserService
}

func NewUserController(service data.UserService) *UserController {
	return &UserController{service: service}
}

func (us *UserController) Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := us.service.Register(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrive the inserted ID"})
		return
	}
	user.ID = oid
	c.JSON(http.StatusCreated, gin.H{"message": "registered successfully", "user": user})

}

func (us *UserController) Login(c *gin.Context) {

	var loginUser models.User
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Printf("login info: %s\n", loginUser)

	token, err := us.service.Login(&loginUser)

	if err != nil || token == "" {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "user logged in successfully", "token": token})

}

func (us *UserController) GetUser(c *gin.Context) {

	role, exists := c.Get("role")

	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "access is for admin-only", "your role": role})
		return
	}

	email := c.Param("email")
	user, err := us.service.GetUser(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (us *UserController) GetUsers(c *gin.Context) {
	role, exists := c.Get("role")

	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "access is for admin-only"})
	}

	users, err := us.service.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)

}

func (us *UserController) DeleteUser(c *gin.Context) {

	role, exists := c.Get("role")

	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"message": "access is for admin-only", "your role": role})
		return
	}

	email := c.Param("email")
	err := us.service.DeleteUser(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfuly"})

}
