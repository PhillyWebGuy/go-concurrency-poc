package main

import (
	"fmt"
	"log"

	"go-concurrency-poc/cmd/http/route"
	"go-concurrency-poc/internal/infrastructure/database"
	"go-concurrency-poc/internal/persistence/book"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	RunHTTP()
}

func RunHTTP() {

	//This calls the mustNewConfiguration function which is expected to return a configuration object.
	//This object is likely to contain settings for the application, such as the HTTP port to listen on.
	config := mustNewConfiguration()

	db := database.MustNewDatabase(
		postgres.Open(
			fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
				config.Database.Host,
				config.Database.User,
				config.Database.Password,
				config.Database.Name,
				config.Database.Port,
			),
		),
		gorm.Config{},
		config.Database.ConnectionRetry,
	)

	b := book.NewService(db)

	h := route.NewHandler(b)

	//This line creates a new Gin engine with the default middleware: logger and recovery (crash-free) middleware.
	r := gin.Default()

	//This defines a route handler for HTTP GET requests to the root ("/") URL. When someone accesses the root URL,
	//the server responds with a JSON object: {"message": "ping pang pung wut up dawg"}.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pang pung wut up dawg",
		})
	})

	route.SetUpRoutes(r, h)

	//This line starts the HTTP server on the port specified in the configuration and listens for requests.
	//If the server fails to start, it will log the error and exit the program.
	//The log.Fatal function is used to print the error message and stop the execution of the program
	//if the r.Run function returns an error.
	log.Fatal(r.Run(config.HTTP.Port))
}
