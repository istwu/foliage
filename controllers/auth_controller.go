package controllers

import (
	"github.com/gofiber/fiber/v2"
	"foliage/services"
)

func Register(c *fiber.Ctx) error {
	return services.RegisterUser(c)
}

func Login(c *fiber.Ctx) error {
	return services.LoginUser(c)
}
