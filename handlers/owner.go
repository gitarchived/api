package handlers

import (
	"github.com/gitarchived/api/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Owner(c *fiber.Ctx, db *gorm.DB) error {
	host := c.Params("host")
	owner := c.Params("owner")

	var results []data.Repository

	if res := db.Where("host = ? AND owner = ?", host, owner).Find(&results); res.Error != nil {
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
		"results": formattedResults,
	})
}
