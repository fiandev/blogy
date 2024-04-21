package router

import (
	"blogy/controllers"
	"blogy/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	// api.Get("/", controllers.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controllers.Login)

	// User
	// user := api.Group("/user")
	// user.Get("/:id", controllers.GetUser)
	// user.Post("/", controllers.CreateUser)
	// user.Patch("/:id", middleware.Protected(), controllers.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), controllers.DeleteUser)

	// post
	// post := api.Group("/posts")
	// post.Get("/", controllers.GetAllposts)
	// post.Get("/:id", controllers.Getpost)
	// post.Post("/", middleware.Protected(), controllers.Createpost)
	// post.Delete("/:id", middleware.Protected(), controllers.Deletepost)
}
