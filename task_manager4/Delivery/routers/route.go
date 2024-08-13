package routers

import (
	"log"
	"task_manager4/Delivery/controllers"
	"task_manager4/Infrastructure"
	"task_manager4/Repositories"
	"task_manager4/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartTaskRoutes(db *mongo.Database, router *gin.Engine) {
	task_repo := Repositories.NewTaskRepo(db, "tasks")
	taskusecase := Usecases.NewTaskUC(task_repo)
	loggedIn := router.Group("")
	loggedIn.Use(Infrastructure.AuthMiddleware)
	loggedIn.GET("/tasks", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.GetAllTasks(taskusecase, c) })
	loggedIn.GET("/tasks/:id", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.GetTaskById(taskusecase, c) })
	loggedIn.POST("/tasks", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.PostTask(taskusecase, c) })
	loggedIn.PUT("/tasks/:id", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.UpdateTask(taskusecase, c) })
	loggedIn.DELETE("/tasks/:id", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.DeleteTask(taskusecase, c) })
}

func StartUserRoutes(db *mongo.Database, router *gin.Engine) {
	user_repo, err := Repositories.NewUserRepo(db, "users")
	if err != nil {
		log.Panic(err.Error())
	}

	pass_s := Infrastructure.PasswordS{}
	token_s := Infrastructure.JwtService{}
	userusecas := Usecases.NewUserUC(user_repo, &pass_s, token_s)

	router.POST("/register", func(c *gin.Context) { controllers.Register(userusecas, c) })
	router.POST("/login", func(c *gin.Context) { controllers.Login(userusecas, c) })

	loggedIn := router.Group("")
	loggedIn.Use(Infrastructure.AuthMiddleware)

	loggedIn.GET("/users", Infrastructure.RoleBasedAuth(true), func(c *gin.Context) { controllers.GetAllUsers(userusecas, c) })
	loggedIn.GET("/users/:id", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.GetUserById(userusecas, c) })
	loggedIn.PUT("/users/:id", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.UpdateUser(userusecas, c) })
	loggedIn.DELETE("/users/:id", Infrastructure.RoleBasedAuth(false), func(c *gin.Context) { controllers.DeleteUser(userusecas, c) })
	loggedIn.PUT("/users/toadmin/:id", Infrastructure.RoleBasedAuth(true), func(c *gin.Context) { controllers.MakeAdmin(userusecas, c) })
}
