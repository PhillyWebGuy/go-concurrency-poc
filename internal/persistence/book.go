package persistence

import (
	"errors"
	"fmt"
	"log/slog"

	"gorm.io/gorm"

	"go-concurrency-poc/internal/domain/entity"
)

func (db Database) ListBooks() ([]entity.Book, error) {
	var books []entity.Book
	err := db.Session.Find(&books).Error
	if err != nil {
		return books, fmt.Errorf("in database.ListBooks: %w", err)
	}

	return books, nil
}

func (db Database) FetchBook(ID int) (entity.Book, error) {
	var user entity.Book
	err := db.Session.Where("ID = ?", ID).First(&book).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Book not found", "ID", ID)
			return book, nil
		}
		return book, fmt.Errorf("in database.FetchBook: %w", err)
	}

	return book, nil
}
