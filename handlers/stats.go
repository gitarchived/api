package handlers

import (
	"github.com/gitarchived/api/data"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Stats(c *fiber.Ctx, db *gorm.DB) error {
	var totalRepos int64

	if result := db.Find(&data.Repository{}).Count(&totalRepos); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching repositories",
		})
	}

	var totalActive int64

	if result := db.Model(&data.Repository{}).Where("deleted = false").Count(&totalActive); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching active repositories",
		})
	}

	totalDeleted := totalRepos - totalActive

	var totalHosts int64

	if result := db.Model(&data.Host{}).Count(&totalHosts); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching hosts",
		})
	}

	var lastRepos []data.Repository

	if result := db.Order("created_at desc").Limit(5).Find(&lastRepos); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error fetching last repositories",
		})
	}

	var formattedLastRepos []data.RepositoryResponse

	for i := range lastRepos {
		formattedLastRepos = append(formattedLastRepos, data.RepositoryResponse{
			ID:         lastRepos[i].ID,
			Host:       lastRepos[i].Host,
			Owner:      lastRepos[i].Owner,
			Name:       lastRepos[i].Name,
			Deleted:    lastRepos[i].Deleted,
			CreatedAt:  lastRepos[i].CreatedAt.Unix(),
			LastCommit: lastRepos[i].LastCommit,
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"repos": fiber.Map{
			"total":   totalRepos,
			"active":  totalActive,
			"deleted": totalDeleted,
			"recent":  formattedLastRepos,
		},
		"hosts": fiber.Map{
			"total": totalHosts,
		},
	})
}
