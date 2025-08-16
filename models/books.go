package models

type Book struct {
    ID       uint   `json:"id" gorm:"primaryKey"`
    Title    string `json:"title"`
    Year     int    `json:"year"`
    AuthorID uint   `json:"author_id"`
}
