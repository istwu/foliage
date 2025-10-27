package controllers

import (
	"foliage/services"

	"github.com/gofiber/fiber/v2"
)

func GetPortfolioByID(c *fiber.Ctx) error {
	return services.GetPortfolioByID(c)
}

func CreatePortfolio(c *fiber.Ctx) error {
	return services.CreatePortfolio(c)
}

func ListPortfolios(c *fiber.Ctx) error {
	return services.ListPortfolios(c)
}

func DeletePortfolio(c *fiber.Ctx) error {
	return services.DeletePortfolio(c)
}

func RenamePortfolio(c *fiber.Ctx) error {
	return services.RenamePortfolio(c)
}
