package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//This creates a new Gin router with default middleware: logger and recovery (crash-free) middleware.
	r := gin.Default()

	//defines a route handler for HTTP GET requests to the root ("/") URL. When someone accesses the root URL,
	//the server responds with a JSON object: {"message": "ping pang pung wut up dawg"}.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping pang pung wut up dawg",
		})
	})

	//This creates a new HTTP server on port 8080 with the Gin router as the handler.
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	//starts the HTTP server in a new goroutine. The server will start serving incoming requests.
	//If the server fails to start, it will panic.
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//This line blocks the program from exiting until it receives a signal.
	//When it receives a signal, it proceeds to the next line.
	<-quit
	//This line prints a message indicating that the server is shutting down.
	println("Shutting down server...")

	//This creates a context that will be canceled after 5 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	//This attempts to gracefully shut down the server within the 5-second timeout. If it fails, it panics.
	if err := srv.Shutdown(ctx); err != nil {
		panic("Server forced to shutdown:")
	}

	//This prints a message indicating that the server has exited.
	println("Server exiting")
}
