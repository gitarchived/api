package main

import (
	"log"
	"os"

	"github.com/gitarchived/api/db"
	"github.com/gitarchived/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
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
	}))

	db := db.Create()

	app.Static("/", "./assets")

	app.Get("/", func(c *fiber.Ctx) error { return handlers.Stats(c, db) })
	app.Post("/create", func(c *fiber.Ctx) error { return handlers.Create(c, db) })
	app.Get("/search", func(c *fiber.Ctx) error { return handlers.Search(c, db) })
	app.Get("/:host/:owner/:name", func(c *fiber.Ctx) error { return handlers.Get(c, db) })
	app.Get("/:host/:owner", func(c *fiber.Ctx) error { return handlers.Owner(c, db) })

	app.Listen(":8080")
}
