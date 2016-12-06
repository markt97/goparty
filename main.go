package main

import "gopkg.in/gin-gonic/gin.v1"

func getter(c *gin.Context) {
	content := gin.H{"Hello": "World"}
	c.JSON(200, content)
}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.GET("/", getter)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
