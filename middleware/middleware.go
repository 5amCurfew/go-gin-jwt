package middleware

import (
	"net/http"

	"github.com/5amCurfew/go-gin-jwt/lib"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AdminMiddleware is a middleware function that handles JWT-based authentication.
// It checks the request for a valid JWT token
// If valid, the request is passed to the next handler
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isValidAdminToken(c) {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
	}
}

// TokenValid checks the validity of the JWT token in the request context.
// It returns an error if the token is invalid or missing.
func isValidAdminToken(c *gin.Context) bool {
	tokenString := lib.ExtractTokenFromRequest(c)
	token, _ := lib.ParseToken(tokenString)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to gather claims"})
		return false
	}
	return claims["adm"].(bool)
}
