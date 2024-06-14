package entity

type Book struct {
	ID    uint   `gorm:"primary_key" json:"id,omitempty"`
	Title string `                   json:"title,omitempty"`
}

//Define interfaces that describe the operations that can be performed on these entities.
//These interfaces will form the contract for your business logic.

type BookRepository interface {
	ListBooks() ([]Book, error)
	FetchBook(ID int) (Book, error)
}
