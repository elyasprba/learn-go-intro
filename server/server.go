package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	run := gin.Default()

	// Setup route untuk Ping
	run.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	// Setup route untuk HealthCheck
	run.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	return run
}
