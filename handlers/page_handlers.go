package handlers

import (
	"go-server/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "application is running"})
}

func RenderIndexPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func RenderForgotPasswordPage(c *gin.Context) {
	c.HTML(http.StatusOK, "forgot_password.html", nil)
}
func RenderHomePage(db *db.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", nil)
	}
}
