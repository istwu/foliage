package services

import (
	"foliage/config"
	"foliage/models"

	"github.com/gofiber/fiber/v2"
)

// Return portfolio info based on ID
func GetPortfolioByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var portfolio models.Portfolio
	if err := config.DB.First(&portfolio, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Portfolio not found"})
	}

	return c.JSON(portfolio)
}

// Add new portfolio info to DB
func CreatePortfolio(c *fiber.Ctx) error {
	data := new(struct {
		ID     uint   `json:"id" gorm:"primaryKey"`
		UserID uint   `json:"user_id"`
		Name   string `json:"name"`
	})

	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if len(data.Name) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Please enter portfolio name"})
	}
	if len(data.Name) > 100 {
		return c.Status(400).JSON(fiber.Map{"error": "Character limit exceeded"})
	}

	portfolio := models.Portfolio{UserID: data.UserID, Name: data.Name}
	config.DB.Create(&portfolio)
	return c.JSON(portfolio)
}

// Get list of specified user's portfolios
func ListPortfolios(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	var portfolios []models.Portfolio
	config.DB.Where("user_id = ?", userID).Find(&portfolios)
	return c.JSON(portfolios)
}

// Delete specified portfolio from DB
func DeletePortfolio(c *fiber.Ctx) error {
	id := c.Params("id")
	var portfolio models.Portfolio
	if err := config.DB.Preload("Posts").First(&portfolio, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Portfolio not found"})
	}

	config.DB.Delete(&models.Post{}, "portfolio_id = ?", portfolio.ID)
	config.DB.Delete(&portfolio)
	return c.JSON(fiber.Map{"message": "Portfolio deleted successfully"})
}

// Change name of portfolio in DB
func RenamePortfolio(c *fiber.Ctx) error {
	id := c.Params("id")
	data := new(struct {
		NewName string `json:"new_name"`
	})
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if len(data.NewName) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Name cannot be blank"})
	}
	if len(data.NewName) > 100 {
		return c.Status(400).JSON(fiber.Map{"error": "Character limit exceeded"})
	}

	var portfolio models.Portfolio
	if err := config.DB.First(&portfolio, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Portfolio not found"})
	}

	portfolio.Name = data.NewName
	config.DB.Save(&portfolio)
	return c.JSON(portfolio)
}
