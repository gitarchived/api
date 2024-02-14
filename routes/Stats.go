package routes

import (
	"github.com/gitarchived/api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Stats(c *fiber.Ctx, db *gorm.DB) error {
	var totalRepos int64

	if result := db.Find(&models.Repository{}).Count(&totalRepos); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching repositories",
		})
	}

	var totalActive int64

	if result := db.Model(&models.Repository{}).Where("deleted = false").Count(&totalActive); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching active repositories",
		})
	}

	totalDeleted := totalRepos - totalActive

	var totalHosts int64

	if result := db.Model(&models.Host{}).Count(&totalHosts); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching hosts",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"repos": fiber.Map{
			"total":   totalRepos,
			"active":  totalActive,
			"deleted": totalDeleted,
		},
		"hosts": fiber.Map{
			"total": totalHosts,
		},
	})
}
