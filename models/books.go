// models/book.go
package models

type Book struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Title    string `json:"title"`
    Year     int    `json:"year"`
    AuthorID uint   `json:"author_id"`
}

// TableName overrides the default table name
func (Book) TableName() string {
    return "books"
}
