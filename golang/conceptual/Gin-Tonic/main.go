package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type TimeoffRequest struct {
	Date   time.Time `form:"date" binding:"required" time_format:"2006-01-02"`
	Amount float64   `form:"amount" binding:"required"`
}

func main() {
	router := gin.Default()
	// static route
	router.GET("/home/", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a static route")
	})
	// Static file method
	router.StaticFile("/static/file", "./go.mod")
	// Static file server
	router.StaticFS("/static/fs", http.Dir("./routing"))

	// Demo appliction
	router.GET("/employee", func(c *gin.Context) {
		c.File("./public/employee.html")
	})

	// router.POST("/employee", func(c *gin.Context) {
	// 	date := c.PostForm("date")
	// 	amount := c.PostForm("amount")
	//
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": "posted",
	// 		"date":   date,
	// 		"amount": amount,
	// 	})
	// })
	// Demo application Parameterised routers
	router.GET("/employee/:username", func(c *gin.Context) {
		employeeName := c.Param("username")
		c.String(http.StatusOK, "This is %s", employeeName)
	})

	// Route group
	{
		admin := router.Group("/admin")
		admin.GET("user")
		admin.GET("roles")
		admin.GET("policies")

	}

	// Demo Request Members
	router.GET("/demo/*rest", func(c *gin.Context) {
		url := c.Request.URL.String()
		headers := c.Request.Header
		cookies := c.Request.Cookies()

		c.IndentedJSON(http.StatusOK, gin.H{
			"url":     url,
			"header":  headers,
			"cookies": cookies,
		})
	})

	// Demo Access simple filed values
	router.GET("/query/*rest", func(c *gin.Context) {
		username := c.Query("username")
		year := c.DefaultQuery("year", strconv.Itoa(time.Now().Year()))
		months := c.QueryArray("months")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"year":     year,
			"months":   months,
		})

	})

	// router.POST("/employee", func(c *gin.Context) {
	// 	amount := c.PostForm("amount")
	// 	date := c.PostFormArray("date")
	//
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"amount": amount,
	// 		"date":   date,
	// 	})
	// })

	router.POST("/employee", func(c *gin.Context) {
		var timeOff TimeoffRequest
		if err := c.ShouldBind(&timeOff); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Piggy not happy with fromat"})
		}
		c.JSON(http.StatusOK, timeOff)

	})

	log.Fatal(router.Run())

}
