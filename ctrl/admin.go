package ctrl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/5amCurfew/go-gin-jwt/lib"
	"github.com/5amCurfew/go-gin-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// ///////////////////////////////////
// GET TOKEN INFO
// ///////////////////////////////////
func AdminToken(c *gin.Context) {
	tokenString := lib.ExtractTokenFromRequest(c)
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing authorization header"})
		return
	}

	token, _ := lib.ParseToken(tokenString)

	var data []byte
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unable to gather claims"})
		return
	}

	userID, _ := strconv.ParseUint(fmt.Sprintf("%.0f", claims["sub"]), 10, 32)
	user, err := models.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, _ = json.Marshal(
		map[string]interface{}{
			"claims":            claims,
			"user":              user,
			"tokenExpirationAt": time.Unix(int64(claims["exp"].(float64)), 0).Format(time.RFC3339),
		},
	)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": json.RawMessage(data)})
}
