package Domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dclaims struct {
	jwt.StandardClaims
	ID      primitive.ObjectID
	Name    string
	Email   string
	IsAdmin bool
}
