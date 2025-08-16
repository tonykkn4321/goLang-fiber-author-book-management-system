package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"

    "goLang-fiber-author-book-management-system/database"
    "goLang-fiber-author-book-management-system/routes"
)

func main() {
    // Load environment-specific .env file
    env := os.Getenv("APP_ENV")
    if env == "" {
        env = "development"
    }

    err := godotenv.Load(".env." + env)
    if err != nil {
        fmt.Printf("Warning: .env.%s file not found, relying on system environment variables\n", env)
    }

    // Create Fiber app
    app := fiber.New()

    // Configure CORS based on environment
    allowedOrigin := os.Getenv("CORS_ORIGIN")
    if allowedOrigin == "" {
        allowedOrigin = "http://localhost:8000"
    }

    app.Use(cors.New(cors.Config{
        AllowOrigins:     allowedOrigin,
        AllowMethods:     "GET,POST,PUT,PATCH,DELETE",
        AllowHeaders:     "Origin, Content-Type, Accept",
        AllowCredentials: true,
    }))

    // Connect to DB
    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    database.DB = db

    // Register routes
    routes.AuthorRoutes(app)
    routes.BookRoutes(app)

    // Start server on configured port
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Printf("Running in %s mode on port %s", env, port)
    log.Fatal(app.Listen(":" + port))
}
