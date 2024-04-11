package ctrl

import (
	"net/http"

	"github.com/5amCurfew/go-gin-jwt/models"
	"github.com/gin-gonic/gin"
)

// ///////////////////////////////////
// REGISTER
// ///////////////////////////////////
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	candidate := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	_, err := candidate.Register()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration successful"})
}
