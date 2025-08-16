// models/author.go
package models

import "gorm.io/gorm"

type Author struct {
    gorm.Model
    Name  string
    Books []Book `gorm:"foreignKey:AuthorID"`
}

// TableName overrides the default table name
func (Author) TableName() string {
    return "authors"
}
