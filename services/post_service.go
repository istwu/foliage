package services

import (
	"fmt"
	"foliage/config"
	"foliage/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateImagePost(c *fiber.Ctx) error {
	data := new(struct {
		UserID      uint   `json:"user_id"`
		PortfolioID uint   `json:"portfolio_id"`
		PostType    string `json:"post_type"`
		ImageURL    string `json:"image_url"`
	})

	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if data.ImageURL == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Image URL is required"})
	}

	post := models.Post{
		UserID:      data.UserID,
		PortfolioID: data.PortfolioID,
		Type:        "image",
		ImageURL:    &data.ImageURL,
	}
	if err := config.DB.Create(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(post)
}

func CreateTextPost(c *fiber.Ctx) error {
	data := new(struct {
		UserID      uint   `json:"user_id"`
		PortfolioID uint   `json:"portfolio_id"`
		PostType    string `json:"post_type"`
		TextBody    string `json:"text_body"`
	})
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if len(data.TextBody) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Post cannot be blank"})
	}
	if len(data.TextBody) > 10000 {
		return c.Status(400).JSON(fiber.Map{"error": "Character limit exceeded"})
	}

	post := models.Post{
		UserID:      data.UserID,
		PortfolioID: data.PortfolioID,
		Type:        data.PostType,
		TextBody:    &data.TextBody,
	}
	if err := config.DB.Create(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(post)
}

func DeletePost(c *fiber.Ctx) error {
	id := c.Params("id")

	// Parse user ID from JSON body (for security verification)
	data := new(struct {
		UserID uint `json:"user_id"`
	})
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Look up the post
	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Post not found"})
	}

	// Verify ownership
	if post.UserID != data.UserID {
		return c.Status(403).JSON(fiber.Map{"error": "Unauthorized: you cannot delete this post"})
	}

	// Proceed to delete
	if err := config.DB.Delete(&post).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete post"})
	}

	return c.JSON(fiber.Map{"message": "Post deleted successfully"})
}

func ListPosts(c *fiber.Ctx) error {
	portfolioIDStr := c.Query("portfolio_id")
	if portfolioIDStr == "" {
		return c.Status(400).JSON(fiber.Map{"error": "portfolio_id is required"})
	}

	portfolioID, err := strconv.ParseUint(portfolioIDStr, 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid portfolio_id"})
	}

	var portfolio models.Portfolio
	// Preload Posts correctly and query by numeric ID
	if err := config.DB.Preload("Posts").First(&portfolio, portfolioID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Portfolio not found"})
	}

	fmt.Printf("%+v\n", portfolio.Posts)
	return c.JSON(portfolio.Posts)
}

func UploadFile(c *fiber.Ctx) error {
	fmt.Println(">>> UploadFile handler reached")

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file uploaded"})
	}
	savePath := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to save file"})
	}
	return c.JSON(fiber.Map{"url": "/" + savePath})
}
