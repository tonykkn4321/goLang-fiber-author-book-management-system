package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "goLang-fiber-author-book-management-system/models"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
    // Use Railway-provided environment variables directly
    host := os.Getenv("PGHOST")
    user := os.Getenv("PGUSER")
    password := os.Getenv("PGPASSWORD")
    dbname := os.Getenv("PGDATABASE")
    port := os.Getenv("PGPORT")

    if host == "" || user == "" || password == "" || dbname == "" || port == "" {
        log.Println("⚠️ Missing one or more required PostgreSQL environment variables")
    }

    // PostgreSQL DSN format
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbname, port,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("❌ failed to connect to database: %w", err)
    }

    // Auto-migrate models
    err = db.AutoMigrate(&models.Author{}, &models.Book{})
    if err != nil {
        return nil, fmt.Errorf("❌ failed to migrate models: %w", err)
    }

    log.Println("✅ Connected to PostgreSQL and migrated models")
    DB = db
    return db, nil
}
