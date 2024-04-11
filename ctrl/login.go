package ctrl

import (
	"net/http"

	"github.com/5amCurfew/go-gin-jwt/models"
	"github.com/gin-gonic/gin"
)

// ///////////////////////////////////
// LOGIN
// ///////////////////////////////////
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	token, err := u.Login()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
