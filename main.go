package main

import (
	"log"
	"os"

	"github.com/gitarchived/api/models"
	"github.com/gitarchived/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	prod := os.Getenv("PRODUCTION")

	if prod == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	app := fiber.New()

	db, err := gorm.Open(postgres.Open(os.Getenv("PG_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// Check if repository table exists
	if !db.Migrator().HasTable(&models.Repository{}) {
		db.Migrator().CreateTable(&models.Repository{})
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  200,
			"message": "OK",
		})
	})

	app.Post("/create", func(c *fiber.Ctx) error { return routes.Create(c, db) })
	app.Get("/search", func(c *fiber.Ctx) error { return routes.Search(c, db) })
	app.Get("/:host/:owner/:name", func(c *fiber.Ctx) error { return routes.Get(c, db) })

	app.Listen(":8080")
}
