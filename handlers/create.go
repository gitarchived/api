package handlers

import (
	"github.com/gitarchived/api/data"
	"github.com/gitarchived/api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CreateRequestBody struct {
	Url string `json:"url" validate:"required,http_url"`
}

func Create(c *fiber.Ctx, db *gorm.DB) error {
	body := new(CreateRequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	validate := validator.New()

	if err := validate.Struct(body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	if !utils.IsUrlOk(body.Url) {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	// Split url
	domain, owner, name := utils.SplitUrl(body.Url)

	if domain == "" || owner == "" || name == "" {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Bad Request",
		})
	}

	var host data.Host

	if result := db.Where("url = ?", "https://"+domain).First(&host); result.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Host not supported",
		})
	}

	if _, err := utils.IsEligible(owner, name); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": err.Error(),
		})
	}

	repository := &data.Repository{
		Host:       host.Name,
		Owner:      owner,
		Name:       name,
		Deleted:    false,
		LastCommit: "",
	}

	if result := db.Select("id").Where("owner = ? AND name = ? AND host = ?", owner, name, host.Name).First(&repository); result.RowsAffected != 0 {
		return c.Status(409).JSON(fiber.Map{
			"status":  409,
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
		"host":    host.Name,
		"owner":   owner,
		"name":    name,
	})
}
