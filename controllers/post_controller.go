package controllers

import (
	"foliage/services"

	"github.com/gofiber/fiber/v2"
)

func CreateImagePost(c *fiber.Ctx) error {
	return services.CreateImagePost(c)
}

func CreateTextPost(c *fiber.Ctx) error {
	return services.CreateTextPost(c)
}

func DeletePost(c *fiber.Ctx) error {
	return services.DeletePost(c)
}

func ListPosts(c *fiber.Ctx) error {
	return services.ListPosts(c)
}

func UploadFile(c *fiber.Ctx) error {
	return services.UploadFile(c)
}
