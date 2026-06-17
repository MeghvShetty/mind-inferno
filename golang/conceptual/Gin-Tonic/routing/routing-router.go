package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World")
	})

	router.POST("/users", func(c *gin.Context) {
		name := c.PostForm("name")
		c.JSON(http.StatusCreated, gin.H{"user": name})
	})

	router.Run(":8080")
}
