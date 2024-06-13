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

	log.Fatal(r.Run(config.HTTP.Port))
}
