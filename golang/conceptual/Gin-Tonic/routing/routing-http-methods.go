package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, greeting)
}

func greeting() string {
	return "hello"
}

func main() {
	router := gin.Default()

	router.GET("/someGet", getting)

	router.Run()
}
