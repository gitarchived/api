package handlers

import (
	"github.com/gitarchived/api/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Search(c *fiber.Ctx, db *gorm.DB) error {
	query := c.Query("q")

	if query == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	var results []data.Repository

	if res := db.Where("LOWER(name) LIKE LOWER(?)", "%"+query+"%").Or("LOWER(owner) LIKE LOWER(?)", "%"+query+"%").Find(&results); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	var formattedResults []data.RepositoryResponse

	for _, result := range results {
		formattedResults = append(formattedResults, data.RepositoryResponse{
			ID:         result.ID,
			Owner:      result.Owner,
			Name:       result.Name,
			Host:       result.Host,
			Deleted:    result.Deleted,
			CreatedAt:  result.CreatedAt.Unix(),
			LastCommit: result.LastCommit,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"results": formattedResults,
	})
}
