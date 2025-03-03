// /utils/jwt.go
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secretkeyjwt")

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(email, role string) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 1).Unix()
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
