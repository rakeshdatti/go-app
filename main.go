// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Go App in Jenkins CI/CD!")
	})

	router.Run(":7080")
}

