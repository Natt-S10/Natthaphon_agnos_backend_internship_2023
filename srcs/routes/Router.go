package routes

import (
	"database/sql"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	r := gin.Default()
	strongPasswdHandler := GetStrongPasswordStepsHandler(db)

	// Group: api
	apiGroup := r.Group("/api")
	{
		apiGroup.POST("/strong_password_steps", strongPasswdHandler)
	}

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}
