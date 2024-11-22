package ctrl

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/5amCurfew/go-gin-jwt/lib"
	"github.com/5amCurfew/go-gin-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ///////////////////////////////////
// GET User information
// ///////////////////////////////////
func GetAdminUser(c *gin.Context) {
	if c.Param("identifier") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id or username required"})
		return
	}

	var data []byte
	var user models.User

	identifier := c.Param("identifier")

	id, err := strconv.Atoi(identifier)

	if err != nil {
		user, err = models.GetUserByUsername(identifier)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, _ = json.Marshal(user)

		c.JSON(http.StatusOK, gin.H{"message": "success", "data": json.RawMessage(data)})
		return
	}

	user, err = models.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ = json.Marshal(user)

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": json.RawMessage(data)})
}

// ///////////////////////////////////
// GET Token information
// ///////////////////////////////////
func GetAdminToken(c *gin.Context) {
	if c.Param("token") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id or username required"})
		return
	}

	token, _ := lib.ParseToken(c.Param("token"))

	var data []byte
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to gather claims"})
		return
	}

	user, err := models.GetUserByID(int(claims["sub"].(float64)))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ = json.Marshal(
		map[string]interface{}{
			"claims":            claims,
			"token":             c.Param("token"),
			"tokenExpirationAt": time.Unix(int64(claims["exp"].(float64)), 0).Format(time.RFC3339),
			"user":              user,
		},
	)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": json.RawMessage(data)})
}
