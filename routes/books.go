package routes

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "goLang-fiber-author-book-management-system/database"
    "goLang-fiber-author-book-management-system/models"
)

func BookRoutes(app *fiber.App) {
    bookGroup := app.Group("/books")

    bookGroup.Get("/", getBooks)
    bookGroup.Get("/:id/", getBookByID)
    bookGroup.Post("/", createBook)
    bookGroup.Put("/:id/", updateBook)
    bookGroup.Patch("/:id/", patchBook)
    bookGroup.Delete("/:id/", deleteBook)
}

func getBooks(c *fiber.Ctx) error {
    var books []models.Book
    database.DB.Preload("Author").Find(&books)
    return c.JSON(books)
}

func getBookByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := database.DB.Preload("Author").First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
    }
    return c.JSON(book)
}

func createBook(c *fiber.Ctx) error {
    fmt.Println("Raw body:", string(c.Body())) // Log raw JSON

    book := new(models.Book)
    if err := c.BodyParser(book); err != nil {
        fmt.Println("BodyParser error:", err) // Log parsing error
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse form data"})
    }

    fmt.Printf("Parsed book: %+v\n", book) // Log parsed struct

    if book.Title == "" || book.Year == 0 || book.AuthorID == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "Title, Year, and AuthorID are required"})
    }

    var author models.Author
    if err := database.DB.First(&author, book.AuthorID).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Author not found"})
    }

    database.DB.Create(&book)
    return c.Status(201).JSON(book)
}


func updateBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
    }

    updated := new(models.Book)
    if err := c.BodyParser(updated); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse form data"})
    }

    if updated.Title == "" || updated.Year == 0 || updated.AuthorID == 0 {
        return c.Status(400).JSON(fiber.Map{"error": "Title, Year, and AuthorID are required"})
    }

    var author models.Author
    if err := database.DB.First(&author, updated.AuthorID).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Author not found"})
    }

    book.Title = updated.Title
    book.Year = updated.Year
    book.AuthorID = updated.AuthorID

    database.DB.Save(&book)
    return c.JSON(book)
}

func patchBook(c *fiber.Ctx) error {
    id := c.Params("id")
    var book models.Book
    if err := database.DB.First(&book, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Book not found"})
    }

    updates := make(map[string]interface{})
    if err := c.BodyParser(&updates); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse form data"})
    }

    if authorID, ok := updates["author_id"].(float64); ok {
        var author models.Author
        if err := database.DB.First(&author, uint(authorID)).Error; err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "Author not found"})
        }
    }

    database.DB.Model(&book).Updates(updates)
    return c.JSON(book)
}

func deleteBook(c *fiber.Ctx) error {
    id := c.Params("id")
    if err := database.DB.Delete(&models.Book{}, id).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to delete book"})
    }
    return c.SendStatus(204)
}
