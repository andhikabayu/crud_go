package middleware

import (
	"crud_go/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := extractBearerToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				c.Abort()
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", &models.User{Username: claims.Username})
	}
}

func extractBearerToken(r *http.Request) (string, error) {
	bearerHeader := r.Header.Get("Authorization")
	if bearerHeader == "" {
		return "", fmt.Errorf("missing Authorization header")
	}

	// Check if the Authorization header has the format "Bearer <token>"
	tokenSplit := strings.Split(bearerHeader, " ")

	if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
		return "", fmt.Errorf("invalid Bearer token format")
	}

	return tokenSplit[1], nil
}
