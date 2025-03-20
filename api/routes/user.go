package routes

import (
	"github.com/deanFL0/blog_api_go/api/handlers"
	"github.com/deanFL0/blog_api_go/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router, service user.Service) {
	users := router.Group("/users")

	users.Get("/", handlers.GetUsers(service))
	users.Get("/:id", handlers.GetUser(service))
	users.Post("/", handlers.AddUser(service))
	users.Put("/:id", handlers.UpdateUser(service))
	users.Delete("/:id", handlers.RemoveUser(service))
}
