package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// External redirect (GET)
	router.GET("/old", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	router.Run(":8080")
}
