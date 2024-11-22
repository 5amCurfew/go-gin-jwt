package lib

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(userID uint, isAdmin bool) (string, error) {
	lifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"iss":   "5am",
		"sub":   userID,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * time.Duration(lifespan)).Unix(),
		"admin": isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))
}

func ParseToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

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
