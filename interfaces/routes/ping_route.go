package routes

import (
	"github.com/gin-gonic/gin"
	"time"
)

func pingRoute(e *gin.Engine) {
	e.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"datetime": time.Now(),
		})
	})
}
