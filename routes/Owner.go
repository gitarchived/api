package routes

import (
	"github.com/gitarchived/api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Owner(c *fiber.Ctx, db *gorm.DB) error {
	host := c.Params("host")
	owner := c.Params("owner")

	var results []models.Repository

	if res := db.Where("host = ? AND owner = ?", host, owner).Find(&results); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"results": results,
	})
}
