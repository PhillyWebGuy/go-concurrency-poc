package entity

type Book struct {
	ID    uint   `gorm:"primary_key" json:"id,omitempty"`
	Title string `                   json:"title,omitempty"`
}
