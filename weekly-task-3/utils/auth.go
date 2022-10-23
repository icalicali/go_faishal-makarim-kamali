package auth

import (
	"log"
	"praktikum/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(userID uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userID,
		"expire": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(constants.JWT_SECRET_KEY))

	if err != nil {
		log.Fatalf("error when creating token: %v \n", tokenString)
	}

	return tokenString
}
