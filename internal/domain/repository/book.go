package repository

import (
	"go-concurrency-poc/internal/domain/entity"
)

//We define the BookRepository interface that will be implemented by the persistence layer.

type BookRepository interface {
	ListBooks() ([]entity.Book, error)
	FetchBook(ID int) (entity.Book, error)
}
