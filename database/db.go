package database

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "goLang-fiber-author-book-management-system/models"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }

    err := godotenv.Load(".env." + env)
    if err != nil {
        fmt.Printf("Warning: .env.%s file not found, relying on system environment variables\n", env)
    }

    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASS")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, pass, host, port, name)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&models.Author{}, &models.Book{})
    if err != nil {
        return nil, err
    }

    DB = db
    return db, nil
}
