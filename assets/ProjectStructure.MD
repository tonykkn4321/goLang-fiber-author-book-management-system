#   Project Structure:
goLang-fiber-author-book-management-system/
├── main.go
├── database/
│   └── db.go
├── models/
│   ├── authors.go
│   └── books.go
├── routes/
│   ├── authors.go
│   └── books.go


#   Initialize the Project:
go mod init goLang-fiber-author-book-management-system
go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/sqlite
go get modernc.org/sqlite
go get github.com/gofiber/fiber/v2/middleware/cors
go mod tidy

#
go run -tags "sqlite_omit_cgo" main.go

#   Optional: Auto-reload with air
#   If you want live reloading during development:
go install github.com/cosmtrek/air@latest
#   Then create an .air.toml config file and run:
air

# The way you've structured your handlers using c.BodyParser(...) in Fiber means they can accept both:
# application/json (standard JSON payloads)
# application/x-www-form-urlencoded (form data from HTML forms or tools like Postman)