package main

import (
	"log"
	"os"

	"github.com/gitarchived/api/models"
	"github.com/gitarchived/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	db, err := gorm.Open(postgres.Open(os.Getenv("PG_URL")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	// Check if repository and host table exists
	if !db.Migrator().HasTable(&models.Repository{}) {
		db.Migrator().CreateTable(&models.Repository{})
	}

	if !db.Migrator().HasTable(&models.Host{}) {
		db.Migrator().CreateTable(&models.Host{})
	}

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error { return routes.Stats(c, db) })
	app.Post("/create", func(c *fiber.Ctx) error { return routes.Create(c, db) })
	app.Get("/search", func(c *fiber.Ctx) error { return routes.Search(c, db) })
	app.Get("/:host/:owner/:name", func(c *fiber.Ctx) error { return routes.Get(c, db) })
	app.Get("/:host/:owner", func(c *fiber.Ctx) error { return routes.Owner(c, db) })

	app.Listen(":8080")
}
