package main

import (
	"github.com/gin-gonic/gin"
	"go_auth/controller"
	"go_auth/middleware"
	"go_auth/models"
	"net/http"
)

func main() {
	models.InitDatabase()
	defer models.DB.Close()

	r := gin.Default()
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/protected", middleware.AuthMiddleware(), ProtectedEndpoint)

	r.Run(":8080")
}

func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a protected endpoint", "user": c.MustGet("username").(string)})
}
