package services

import (
	// "errors"
	"foliage/config"
	"foliage/models"

	"github.com/gofiber/fiber/v2"
)

// Add new user to database
func RegisterUser(c *fiber.Ctx) error {
	data := new(struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		// Password string `json:"password"`
	})
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if len(data.Username) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Username cannot be blank"})
	}

	// hashed, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	user := models.User{Username: data.Username} // need to include email and password hash here later
	if err := config.DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Username already exists"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Registration successful",
		"user":    user,
	})
}

// Check username against DB
func LoginUser(c *fiber.Ctx) error {
	data := new(struct {
		Username string `json:"username"`
		// Password string `json:"password"`
	})
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Log the user in if they exist in the DB
	var user models.User
	if err := config.DB.Where("username = ?", data.Username).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid username"})
	}

	// if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(data.Password)) != nil {
	// 	return c.Status(400).JSON(fiber.Map{"error": "Password is incorrect"})
	// }

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}
