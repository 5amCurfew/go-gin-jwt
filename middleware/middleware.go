package middleware

import (
	"net/http"

	"github.com/5amCurfew/go-gin-jwt/lib"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ///////////////////////////////////
// /public/* middleware
// ///////////////////////////////////
func PublicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

// ///////////////////////////////////
// /admin/* middleware
// ///////////////////////////////////
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isValidAdminToken(c) {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized token"})
			c.Abort()
			return
		}
	}
}

func isValidAdminToken(c *gin.Context) bool {
	tokenString := lib.ExtractTokenFromRequest(c)
	token, _ := lib.ParseToken(tokenString)
	if token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		return claims["admin"].(bool)
	}

	return false
}
