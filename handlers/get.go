package handlers

import (
	"github.com/gitarchived/api/data"
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

	var results data.Repository

	if res := db.Where("host = ? AND owner = ? AND name = ?", host, owner, name).First(&results); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	formattedResults := data.RepositoryResponse{
		ID:         results.ID,
		Host:       results.Host,
		Owner:      results.Owner,
		Name:       results.Name,
		Deleted:    results.Deleted,
		CreatedAt:  results.CreatedAt.Unix(),
		LastCommit: results.LastCommit,
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"data":    formattedResults,
	})
}
