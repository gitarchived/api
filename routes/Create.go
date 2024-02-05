package routes

import (
	"github.com/gitarchived/api/models"
	"github.com/gitarchived/api/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateRequestBody struct {
	Host  string `json:"host"` // Only github supported for now
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

func Create(c *fiber.Ctx, db *gorm.DB) error {
	body := new(CreateRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	if body.Host == "" || body.Owner == "" || body.Name == "" {
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

	if _, err := utils.IsEligible(body.Owner, body.Name); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	repository := &models.Repository{
		Host:       body.Host,
		Owner:      body.Owner,
		Name:       body.Name,
		Deleted:    false,
		LastCommit: "",
	}

	if result := db.Select("id").Where("owner = ? AND name = ? AND host = ?", body.Owner, body.Name, body.Host).First(&repository); result.RowsAffected != 0 {
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
		"owner":   body.Owner,
		"name":    body.Name,
	})
}
