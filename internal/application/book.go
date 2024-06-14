package application

import (
	"go-concurrency-poc/internal/domain/entity"
)

//In your application logic, you should only depend on the interfaces defined in the domain layer,
//not the concrete implementations. This way, your business logic is decoupled from the infrastructure
//layer and you can easily swap out implementations if needed.

type BookService struct {
	repo entity.BookRepository
}

func NewBookService(repo entity.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) GetBookByID(id int) (entity.Book, error) {
	return s.repo.FetchBook(id)
}

func (s *BookService) GetAllBooks() ([]entity.Book, error) {
	return s.repo.ListBooks()
}
