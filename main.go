package main

import (
	"github.com/gin-gonic/gin"
	"go_auth/internal/domain/model"
	"go_auth/internal/interface/http/handler"
	"go_auth/internal/interface/http/middleware"
	"net/http"
)

func main() {
	model.InitDatabase()
	defer model.DB.Close()

	r := gin.Default()
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.GET("/protected", middleware.AuthMiddleware(), ProtectedEndpoint)

	r.Run(":8080")
}

func ProtectedEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a protected endpoint", "user": c.MustGet("username").(string)})
}
