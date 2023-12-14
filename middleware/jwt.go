package middleware

import (
	"crud_go/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("thisIsTheSecretKeyYeeha")

// Claims represents the claims carried by the JWT.
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for a user.
func GenerateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
