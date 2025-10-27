package routes

import (
	"foliage/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	portfolio := api.Group("/portfolio")
	portfolio.Get("/posts", controllers.ListPosts)
	portfolio.Get("/:id", controllers.GetPortfolioByID)
	portfolio.Post("/", controllers.CreatePortfolio)
	portfolio.Get("/", controllers.ListPortfolios)
	portfolio.Delete("/:id", controllers.DeletePortfolio)
	portfolio.Put("/:id/rename", controllers.RenamePortfolio)

	post := api.Group("/post")
	post.Post("/upload", controllers.UploadFile)
	post.Post("/image", controllers.CreateImagePost)
	post.Post("/text", controllers.CreateTextPost)
	post.Delete("/:id", controllers.DeletePost)
}
