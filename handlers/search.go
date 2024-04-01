package handlers

import (
	"github.com/gitarchived/api/data"
	"github.com/gitarchived/api/utils"
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

	var formattedQuery = utils.ExamineQuery(query)
	var results []data.Repository

	dbQuery := db.Model(&data.Repository{})

	if formattedQuery.Owner == formattedQuery.Name {
		dbQuery = dbQuery.Where("LOWER(owner) LIKE LOWER(?)", "%"+formattedQuery.Owner+"%")
		dbQuery = dbQuery.Or("LOWER(name) LIKE LOWER(?)", "%"+formattedQuery.Name+"%")
	} else {
		dbQuery = dbQuery.Where("owner = ? AND LOWER(name) LIKE LOWER(?)", formattedQuery.Owner, "%"+formattedQuery.Name+"%")
	}

	if res := dbQuery.Find(&results); res.Error != nil {
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
