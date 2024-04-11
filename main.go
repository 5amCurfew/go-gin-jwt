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

	models.Connect()
	r := gin.Default()
	public := r.Group("/api")

	public.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong 🏓"})
	})

	// curl -X POST localhost:8080/api/register -H "Content-Type: application/json" -d '{"username": "<USERNAME>", "password": "<PASSWORD>"}'
	public.POST("/register", ctrl.Register)
	// curl -X POST localhost:8080/api/login -H "Content-Type: application/json" -d '{"username": "<USERNAME>", "password": "<PASSWORD>"}'
	public.POST("/login", ctrl.Login)

	protected := r.Group("/api/admin")
	protected.Use(middleware.JwtAuthMiddleware())
	// curl -X GET localhost:8080/api/admin/token -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: bearer <TOKEN>" | jq .
	protected.GET("/token", ctrl.GetTokenInfo)

	r.Run(":8080")
}
