package routes

import (
	"github.com/gitarchived/api/models"
	"github.com/gitarchived/api/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateRequestBody struct {
	Host string `json:"host"` // Only github supported for now
	Repo string `json:"repo"`
}

func Create(c *fiber.Ctx, db *gorm.DB) error {
	body := new(CreateRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	if body.Host == "" || body.Repo == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	if body.Host != "github" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Host not supported",
		})
	}

	if exist := utils.RepoExist(body.Repo); !exist {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Repository not found",
		})
	}

	lastCommit, err := utils.GetLastCommit(body.Repo)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	repository := &models.Repository{
		Name:       body.Repo,
		Host:       body.Host,
		LastCommit: lastCommit,
	}

	if result := db.Select("id").Where("name = ? AND host = ?", body.Repo, body.Host).First(&repository); result.RowsAffected != 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Repository already added",
		})
	}

	result := db.Create(&repository)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Repository added successfully",
		"host":    body.Host,
		"repo":    body.Repo,
	})
}
