package routes

import (
    "github.com/gofiber/fiber/v2"
    "goLang-fiber-author-book-management-system/database"
    "goLang-fiber-author-book-management-system/models"
)

func AuthorRoutes(app *fiber.App) {
    authorGroup := app.Group("/authors")

    authorGroup.Get("/", getAuthors)
    authorGroup.Get("/:id/", getAuthorByID)
    authorGroup.Post("/", createAuthor)
    authorGroup.Put("/:id/", updateAuthor)
    authorGroup.Patch("/:id/", patchAuthor)
    authorGroup.Delete("/:id/", deleteAuthor)
}

func getAuthors(c *fiber.Ctx) error {
    var authors []models.Author
    database.DB.Preload("Books").Find(&authors)
    return c.JSON(authors)
}

func getAuthorByID(c *fiber.Ctx) error {
    id := c.Params("id")
    var author models.Author
    if err := database.DB.Preload("Books").First(&author, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Author not found"})
    }
    return c.JSON(author)
}

func createAuthor(c *fiber.Ctx) error {
    author := new(models.Author)
    if err := c.BodyParser(author); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse form data"})
    }

    if author.FirstName == "" || author.LastName == "" {
        return c.Status(400).JSON(fiber.Map{"error": "First name and last name are required"})
    }

    database.DB.Create(&author)
    return c.Status(201).JSON(author)
}

func updateAuthor(c *fiber.Ctx) error {
    id := c.Params("id")
    var author models.Author
    if err := database.DB.First(&author, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Author not found"})
    }

    updated := new(models.Author)
    if err := c.BodyParser(updated); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse form data"})
    }

    if updated.FirstName == "" || updated.LastName == "" {
        return c.Status(400).JSON(fiber.Map{"error": "First name and last name are required"})
    }

    author.FirstName = updated.FirstName
    author.LastName = updated.LastName

    database.DB.Save(&author)
    return c.JSON(author)
}

func patchAuthor(c *fiber.Ctx) error {
    id := c.Params("id")
    var author models.Author
    if err := database.DB.First(&author, id).Error; err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Author not found"})
    }

    updates := make(map[string]interface{})
    if err := c.BodyParser(&updates); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse form data"})
    }

    database.DB.Model(&author).Updates(updates)
    return c.JSON(author)
}

func deleteAuthor(c *fiber.Ctx) error {
    id := c.Params("id")
    if err := database.DB.Delete(&models.Author{}, id).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Failed to delete author"})
    }
    return c.SendStatus(204)
}
