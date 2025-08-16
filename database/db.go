package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "goLang-fiber-author-book-management-system/models"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Auto-migrate your models
    err = db.AutoMigrate(&models.Author{}, &models.Book{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
