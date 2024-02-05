package routes

import (
	"time"

	"github.com/gitarchived/api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Get(c *fiber.Ctx, db *gorm.DB) error {
	host := c.Params("host")
	owner := c.Params("owner")
	name := c.Params("name")

	if owner == "" || name == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	var results models.Repository

	data := db.Where("host = ? AND owner = ? AND name = ?", host, owner, name).First(&results)

	if data.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	formattedResults := models.RepositoryResponse{
		ID:         results.ID,
		Host:       results.Host,
		Owner:      results.Owner,
		Name:       results.Name,
		Deleted:    results.Deleted,
		CreatedAt:  results.CreatedAt.Format(time.RFC3339),
		LastCommit: results.LastCommit,
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"data":    formattedResults,
	})
}
