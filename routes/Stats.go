package routes

import (
	"github.com/gitarchived/api/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Stats(c *fiber.Ctx, db *gorm.DB) error {
	totalRepos := db.Model(&models.Repository{}).Find(&models.Repository{}).RowsAffected
	totalHosts := db.Model(&models.Host{}).Find(&models.Host{}).RowsAffected
	totalDeleted := db.Model(&models.Repository{}).Where("deleted = ?", true).Find(&models.Repository{}).RowsAffected
	totalActive := totalRepos - totalDeleted

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
