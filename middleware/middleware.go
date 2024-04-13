package middleware

import (
	"net/http"

	lib "github.com/5amCurfew/go-gin-jwt/lib"
	"github.com/gin-gonic/gin"
)

// JwtAuthMiddleware is a middleware function that handles JWT-based authentication.
// It checks the request for a valid JWT token
// If valid, the request is passed to the next handler
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// TokenValid checks the validity of the JWT token in the request context.
// It returns an error if the token is invalid or missing.
func TokenValid(c *gin.Context) error {
	tokenString := lib.ExtractTokenFromRequest(c)
	_, err := lib.ParseToken(tokenString)
	if err != nil {
		return err
	}
	return nil
}
