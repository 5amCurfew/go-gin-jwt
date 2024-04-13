package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// GenerateToken generates a JWT token for the given user ID.
// userID The ID of the user to generate the token for and returns the generated JWT token as a string, or an error if the token could not be generated.
func GenerateToken(userID uint) (string, error) {
	lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"aut": true,
		"sub": userID,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * time.Duration(lifespan)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

// ParseToken parses the given JWT token string and returns the parsed token and any error that occurred during parsing.
func ParseToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

// ExtractTokenFromRequest extracts the JWT token from the request context.
// It expects the token to be present in the "Authorization" header, with the format "Bearer <token>".
// If the token is not present or the format is invalid, an empty string is returned.
func ExtractTokenFromRequest(c *gin.Context) string {
	if c.Request.Header.Get("Authorization") == "" {
		return ""
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
