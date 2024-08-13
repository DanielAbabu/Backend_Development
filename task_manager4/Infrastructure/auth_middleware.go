package Infrastructure

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func AuthMiddleware(c *gin.Context) {

	// check for the authorization
	auth := c.GetHeader("Authorization")

	if auth == "" {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"msg": "NOT AUTHORIZED"})
		c.Abort()
		return
	}

	JW := JwtService{}
	err := JW.TokenValidate(auth)

	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// move to the next step
	c.Next()
}
