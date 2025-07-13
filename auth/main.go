package main

import (
	"os"

	"github.com/archon42x/agora/auth/handler"
	"github.com/archon42x/agora/common/jwt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/register", handler.RegisterHandler)
	r.POST("/login", handler.LoginHandler)

	protected := r.Group("/", jwt.AuthMiddleware())
	protected.GET("/username", handler.UsernameHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	r.Run(":" + port)
}
