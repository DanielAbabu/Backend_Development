package Infrastructure

import (
	"errors"
	"fmt"
	"strings"
	"task_manager5/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService struct{}

func (JwtService) TokenValidate(auth string) error {
	authSplit := strings.Split(auth, " ")

	if len(authSplit) != 2 || strings.ToLower(authSplit[0]) != "bearer" {
		return errors.New("not authorized")
	}

	token, err := jwt.Parse(authSplit[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return errors.New("invalid jwt")
	}

	return nil
}

func (JwtService) CreateToken(usr Domain.UserInput) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&Domain.Dclaims{
			ID:             usr.ID,
			Name:           usr.Name,
			Email:          usr.Email,
			IsAdmin:        usr.IsAdmin,
			StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour * 24).Unix()},
		})

	jwtToken, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
