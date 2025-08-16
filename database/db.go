package database

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "goLang-fiber-author-book-management-system/models"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
    // Replace with your actual MySQL credentials
    username := "root"
    password := "Aa161616"
    host := "127.0.0.1"
    port := "3306"
    dbname := "author_book_management_system"

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        username, password, host, port, dbname)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
