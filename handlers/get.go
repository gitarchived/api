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

	var result data.Repository

	if res := db.Where("host = ? AND owner = ? AND name = ?", host, owner, name).First(&result); res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	formattedResult := data.RepositoryResponse{
		ID:         result.ID,
		Host:       result.Host,
		Owner:      result.Owner,
		Name:       result.Name,
		Deleted:    result.Deleted,
		CreatedAt:  result.CreatedAt.Unix(),
		LastCommit: result.LastCommit,
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "OK",
		"data":    formattedResult,
	})
}
