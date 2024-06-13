package book

import (
	"fmt"

	"go-concurrency-poc/internal/domain/entity"
)

type databaseSession interface {
	ListBooks() ([]entity.Book, error)
	FetchBook(ID int) (entity.Book, error)
}

type Service struct {
	Database databaseSession
}

func NewService(db databaseSession) Service {
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
