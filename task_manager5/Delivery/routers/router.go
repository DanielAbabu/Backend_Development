package routers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Run(db *mongo.Database) {
	router := gin.Default()
	StartTaskRoutes(db, router)
	StartUserRoutes(db, router)

	router.Run("localhost:8080")
}
