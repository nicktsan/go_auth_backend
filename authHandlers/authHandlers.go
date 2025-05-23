package authHandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	fmt.Println("Login handler called")
	c.SetCookie("refreshTokenString", "sample_refresh_token", 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": "sample_access_token",
	})
}

func IsAuthenticated(c *gin.Context) {
	cookie, err := c.Cookie("refreshTokenString")
	fmt.Println(cookie)
	if err != nil {
		c.String(http.StatusNotFound, "Cookie not found")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cookie_value": cookie,
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
