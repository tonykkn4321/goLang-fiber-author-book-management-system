package database

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"

    "goLang-fiber-author-book-management-system/config"
    "goLang-fiber-author-book-management-system/models"
)

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
    var dsn string
    var dialector gorm.Dialector

    switch config.AppEnv {
    case "development":
        dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
            config.DB.Username,
            config.DB.Password,
            config.DB.Host,
            config.DB.Port,
            config.DB.Database,
        )
        dialector = mysql.Open(dsn)

    case "test":
        dialector = sqlite.Open(config.DB.Storage)

    case "production":
        dsn = os.Getenv(config.DB.UseEnvVariable)
        if dsn == "" {
            return nil, fmt.Errorf("❌ DATABASE_URL not set")
        }
        dialector = postgres.Open(dsn)

    default:
        return nil, fmt.Errorf("❌ Unknown environment: %s", config.AppEnv)
    }

    db, err := gorm.Open(dialector, &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("❌ failed to connect to database: %w", err)
    }

    err = db.AutoMigrate(&models.Author{}, &models.Book{})
    if err != nil {
        return nil, fmt.Errorf("❌ failed to migrate models: %w", err)
    }

    log.Println("✅ Connected and migrated models")
    return db, nil
}
