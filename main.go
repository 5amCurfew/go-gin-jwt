package main

import (
	"net/http"

	"github.com/5amCurfew/go-gin-jwt/ctrl"
	"github.com/5amCurfew/go-gin-jwt/middleware"
	"github.com/5amCurfew/go-gin-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load(".env")
	if err != nil {
		panic("error parsing .env")
	}

	// Connect to auth database
	models.ConnectToAuthDatabase()
	r := gin.Default()

	public := r.Group("/")
	public.Use(middleware.PublicMiddleware())
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong 🏓"})
	})

	auth := r.Group("/auth")
	// curl -X POST localhost:8080/auth/register -H "Content-Type: application/json" -d '{"username": "<USERNAME>", "password": "<PASSWORD>"}'
	auth.POST("/register", ctrl.Register)
	// curl -X POST localhost:8080/auth/login -H "Content-Type: application/json" -d '{"username": "<USERNAME>", "password": "<PASSWORD>"}'
	auth.POST("/login", ctrl.Login)

	admin := r.Group("/admin")
	admin.Use(middleware.AdminMiddleware())
	// curl -X GET localhost:8080/admin/token -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: bearer <TOKEN>" | jq .
	admin.GET("/token", ctrl.AdminToken)

	r.Run(":8080")
}
