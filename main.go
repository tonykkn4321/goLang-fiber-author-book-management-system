package main

import (
    "fmt"
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/joho/godotenv"

    "goLang-fiber-author-book-management-system/config"
    "goLang-fiber-author-book-management-system/database"
    "goLang-fiber-author-book-management-system/routes"
)

func main() {
    // Load .env file
    _ = godotenv.Load(".env." + config.AppEnv)

    // Load config
    config.LoadConfig()

    // Create Fiber app
    app := fiber.New()

    // CORS setup
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
        log.Fatalf("❌ Failed to connect to database: %v", err)
    }
    database.DB = db
    fmt.Println("✅ Connected to database")

    // Health check
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "ok",
            "env":    config.AppEnv,
        })
    })

    // Routes
    routes.AuthorRoutes(app)
    routes.BookRoutes(app)

    // Start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }

    log.Printf("🚀 Running in %s mode on port %s", config.AppEnv, port)
    log.Fatal(app.Listen(":" + port))
}
