package routes

import (
	"foliage/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Base API group
	api := app.Group("/api")

	// Auth routes
	api.Post("/auth/register", controllers.Register)
	api.Post("/auth/login", controllers.Login)

	// Portfolio routes
	api.Get("/portfolio/posts", controllers.ListPosts)
	api.Get("/portfolio/:id", controllers.GetPortfolioByID)
	api.Post("/portfolio", controllers.CreatePortfolio)
	api.Get("/portfolio", controllers.ListPortfolios)
	api.Delete("/portfolio/:id", controllers.DeletePortfolio)
	api.Put("/portfolio/:id/rename", controllers.RenamePortfolio)

	// Post routes
	api.Post("/post/upload", controllers.UploadFile)
	api.Post("/post/image", controllers.CreateImagePost)
	api.Post("/post/text", controllers.CreateTextPost)
	api.Post("/post/reorder", controllers.ReorderPosts)
	api.Delete("/post/:id", controllers.DeletePost)
}
