package book

import (
	"fmt"

	"go-concurrency-poc/internal/domain/entity"
	"go-concurrency-poc/internal/domain/repository"
)

type Service struct {
	Database repository.BookRepository
}

func NewService(db repository.BookRepository) Service {
	return Service{Database: db}
}

func (s Service) List() ([]entity.Book, error) {
	books, err := s.Database.ListBooks()
	if err != nil {
		return []entity.Book{}, fmt.Errorf("in book.List: %w", err)
	}
	return books, nil
}

func (s Service) Fetch(ID int) (entity.Book, error) {
	book, err := s.Database.FetchBook(ID)
	if err != nil {
		return entity.Book{}, fmt.Errorf("in book.Fetch: %w", err)
	}
	return book, nil
}
