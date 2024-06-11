package main

import (
	"os"
	"v2/internal/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {

	//This creates a new Gin router with default middleware: logger and recovery (crash-free) middleware.
	r := gin.Default()

	//defines a route handler for HTTP GET requests to the root ("/") URL. When someone accesses the root URL,
	//the server responds with a JSON object: {"message": "ping pang pung wut up dawg"}.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "wut up dawg",
		})
	})

	//This line starts the server on port env port. The server will listen for incoming requests and respond to them.
	r.Run(":" + os.Getenv("PORT"))

}
