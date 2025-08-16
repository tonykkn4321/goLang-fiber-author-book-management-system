package models

type Author struct {
    ID        uint   `json:"id" gorm:"primaryKey"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Books     []Book `json:"books" gorm:"foreignKey:AuthorID"`
}
