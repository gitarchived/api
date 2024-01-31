package main

import (
	"log"
	"os"

	"github.com/gitarchived/api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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

	app.Listen(":8080")
}
