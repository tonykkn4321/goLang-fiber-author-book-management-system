package config

import (
    "fmt"
    "os"
)

type DBConfig struct {
    Dialect         string
    Username        string
    Password        string
    Database        string
    Host            string
    Port            string
    Storage         string // for SQLite
    UseEnvVariable  string
    SSLRequired     bool
    Logging         bool
}

var AppEnv string
var DB DBConfig

func LoadConfig() {
    AppEnv = os.Getenv("APP_ENV")
    if AppEnv == "" {
        AppEnv = "development"
    }

    switch AppEnv {
    case "development":
        DB = DBConfig{
            Dialect:  "mysql",
            Username: os.Getenv("DB_USER"),
            Password: os.Getenv("DB_PASS"),
            Database: os.Getenv("DB_NAME"),
            Host:     os.Getenv("DB_HOST"),
            Port:     os.Getenv("DB_PORT"),
            Logging:  true,
        }
    case "test":
        DB = DBConfig{
            Dialect: "sqlite",
            Storage: "./test.sqlite",
            Logging: false,
        }
    case "production":
        DB = DBConfig{
            Dialect:        "postgres",
            UseEnvVariable: "DATABASE_URL",
            SSLRequired:    true,
            Logging:        false,
        }
    default:
        fmt.Printf("⚠️ Unknown APP_ENV: %s\n", AppEnv)
    }
}
