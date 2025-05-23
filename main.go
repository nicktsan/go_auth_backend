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
	r.GET("/login", authHandlers.HandleLogin)
	r.GET("/isAuthenticated", authHandlers.IsAuthenticated)

	r.Run(":8080")
}
