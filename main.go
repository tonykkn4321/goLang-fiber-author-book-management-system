package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "goLang-fiber-author-book-management-system/database"
    "goLang-fiber-author-book-management-system/routes"
)

func main() {
    app := fiber.New()

    // Enable CORS for frontend access
    app.Use(cors.New(cors.Config{
        AllowOrigins:     "http://localhost:8000",           
        AllowMethods:     "GET,POST,PUT,PATCH,DELETE",            
        AllowHeaders:     "Origin, Content-Type, Accept",    
        AllowCredentials: true,
    }))


    // Connect to DB and assign to global variable
    db, err := database.ConnectDB()
    if err != nil {
        panic("Failed to connect to database")
    }
    database.DB = db

    // Register routes
    routes.AuthorRoutes(app)
    routes.BookRoutes(app)

    app.Listen(":3000")
}
