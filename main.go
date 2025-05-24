package main

import (
	"auth_backend/authHandlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Use the CORS middleware defined below
	r.Use(authHandlers.CORSMiddleware())

	// Route for generating tokens
	r.POST("/login", authHandlers.HandleLogin)
	r.POST("/refreshToken", authHandlers.HandleRefresh)

	r.Run("localhost:8080")
}
