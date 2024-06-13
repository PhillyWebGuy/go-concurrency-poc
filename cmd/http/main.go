package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	RunHTTP()
}

func RunHTTP() {
	config := mustNewConfiguration()

	r := gin.Default()

	//This defines a route handler for HTTP GET requests to the root ("/") URL. When someone accesses the root URL,
	//the server responds with a JSON object: {"message": "ping pang pung wut up dawg"}.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pang pung wut up dawg",
		})
	})

	log.Fatal(r.Run(config.HTTP.Port))
}
