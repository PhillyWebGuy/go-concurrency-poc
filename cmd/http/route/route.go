package route

import (
	"github.com/gin-gonic/gin"

	"go-concurrency-poc/internal/domain/entity"
)

type responseMessage struct {
	Message string `json:"message"`
}

type responseID struct {
	ObjectID int `json:"object_id"`
}

type responseError struct {
	Error string `json:"error"`
}

type bookService interface {
	List() ([]entity.Book, error)
	Fetch(ID int) (entity.Book, error)
}

type Handler struct {
	bookService bookService
}

func NewHandler(bookService bookService) Handler {
	return Handler{bookService: bookService}
}

func SetUpRoutes(r *gin.Engine, h Handler) {
	healthCheck(r.Group("/health-check"), h)

	bookRoutes(r.Group("/book"), h)

	notFound(r)
}
