package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		// Authentication
		auth := api.Group("/auth")
		{
			auth.POST("/login", loginHandler)
			auth.POST("/register", registerHandler)
			//auth.POST("passwordReset", passwordrest)
		}

		// Bundles
		bundles := api.Group("/bundles")
		{
			// Create a bundle
			bundles.POST("", createBundleHandler)

			// Save a bundle
			bundles.POST("/save", saveBundleHandler)

			// Get previous bundles
			// Optional: ?category=adventure|dining|sports|music
			bundles.GET("/previous", previousBundlesHandler)
		}

		// Profile
		profile := api.Group("/profile")
		{
			profile.GET("", getProfileHandler)
			profile.PUT("", updateProfileHandler)
		}

		// Feedback
		api.POST("/feedback", feedbackHandler)

		// Explore page
		api.GET("/explore", exploreHandler)
	}

	router.Run(":8080")
}

// -------------------------
// Handler Functions
// -------------------------

func loginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login API",
	})
}

func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration API",
	})
}

func createBundleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Create Bundle API",
	})
}

func saveBundleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Save Bundle API",
	})
}

func previousBundlesHandler(c *gin.Context) {
	category := c.Query("category")

	if category == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Returning all previous bundles",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Returning filtered bundles",
		"category": category,
	})
}

func getProfileHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Profile API",
	})
}

func updateProfileHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update Profile API",
	})
}

func feedbackHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Feedback API",
	})
}

func exploreHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Explore API",
	})
}
