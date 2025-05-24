package authHandlers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	//Pretend to validate user credentials
	// In a real application, you would validate the user credentials here
	c.SetCookie("refresh_token", "sample_refresh_token", 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{
		"access_token": "sample_access_token",
	})
}

func HandleLogout(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.String(http.StatusUnauthorized, "Refresh token not found.")
		return
	}
	values := []string{"sample_refresh_token", "sample_refresh_token_2"}
	pattern := strings.Join(values, "|")
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllString(cookie, -1)
	// Invalidate the refresh token by setting it to an empty value
	// In a real application, you would also invalidate the token in your database or cache
	if len(matches) > 0 {
		c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "Logged out successfully",
		})
	} else {
		c.String(http.StatusUnauthorized, "Invalid refresh token.")
	}

}

func HandleRefresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.String(http.StatusUnauthorized, "Refresh token not found.")
		return
	}
	values := []string{"sample_refresh_token", "sample_refresh_token_2"}
	pattern := strings.Join(values, "|")
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllString(cookie, -1)
	// Here you would typically validate the refresh token and issue a new access token
	if len(matches) > 0 {
		c.SetCookie("refresh_token", "sample_refresh_token_2", 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{
			"access_token": "sample_access_token",
		})
	} else {
		c.String(http.StatusUnauthorized, "Invalid refresh token")
	}
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
