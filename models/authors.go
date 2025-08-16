// models/author.go
package models

type Author struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Books     []Book `json:"books" gorm:"foreignKey:AuthorID"`
}


// TableName overrides the default table name
func (Author) TableName() string {
    return "authors"
}
