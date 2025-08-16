// models/book.go
package models

import "gorm.io/gorm"

type Book struct {
    gorm.Model
    Title    string
    AuthorID uint
}

// TableName overrides the default table name
func (Book) TableName() string {
    return "books"
}
