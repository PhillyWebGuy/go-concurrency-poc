package route

import (
	"log/slog"
	"net/http"
	"strconv"

	"go-concurrency-poc/internal/domain/entity"

	"github.com/gin-gonic/gin"
)

type responseOneBook struct {
	Book entity.Book `json:"book"`
}

type responseAllBooks struct {
	Books []entity.Book `json:"books"`
}

func bookRoutes(r *gin.RouterGroup, h Handler) {
	// list all books
	r.GET("/", func(c *gin.Context) {
		// get values from db
		books, err := h.bookService.List()
		if err != nil {
			slog.Error("error getting all locations", "error", err)
			c.JSON(http.StatusInternalServerError, responseError{Error: "Error retrieving data"})
			return
		}

		// return response
		c.JSON(http.StatusOK, responseAllBooks{Books: books})
	})

	// fetch a book by ID
	r.GET("/:ID", func(c *gin.Context) {
		// get and validate ID
		idString := c.Param("ID")
		ID, err := strconv.Atoi(idString)
		if err != nil {
			slog.Error("error getting ID", "error", err)
			c.JSON(http.StatusBadRequest, responseError{Error: "Not a valid ID"})
			return
		}

		// get values from db
		book, err := h.bookService.Fetch(ID)
		if err != nil {
			slog.Error("error getting all locations", "error", err)
			c.JSON(http.StatusInternalServerError, responseError{Error: "Error retrieving data"})
			return
		}

		// return response
		c.JSON(http.StatusOK, responseOneBook{Book: book})
	})
}
