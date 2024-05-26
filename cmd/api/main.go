package main

import (
	"log"
	"os"
	"time"

	"github.com/gitarchived/api/db"
	"github.com/gitarchived/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        20,
		Expiration: 20 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{
				"message": "Too many requests",
				"status":  429,
			})
		},
	}))

	db := db.Create()

	app.Static("/", "./assets")

	app.Get("/", func(c *fiber.Ctx) error { return handlers.Stats(c, db) })
	app.Post("/create", func(c *fiber.Ctx) error { return handlers.Create(c, db) })
	app.Get("/search", func(c *fiber.Ctx) error { return handlers.Search(c, db) })
	app.Get("/:host/:owner/:name", func(c *fiber.Ctx) error { return handlers.Get(c, db) })
	app.Get("/:host/:owner", func(c *fiber.Ctx) error { return handlers.Owner(c, db) })

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not Found",
			"status":  404,
		})
	})

	app.Listen(":8080")
}
