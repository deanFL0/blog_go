package routes

import (
	"github.com/deanFL0/blog_api_go/api/handlers"
	"github.com/deanFL0/blog_api_go/pkg/article"
	"github.com/gofiber/fiber/v2"
)

func ArticleRouter(router fiber.Router, service article.Service) {
	articles := router.Group("/articles")

	articles.Get("/", handlers.GetArticles(service))
	articles.Get("/:id", handlers.GetArticle(service))
	articles.Post("/", handlers.AddArticle(service))
	articles.Put("/:id", handlers.UpdateArticle(service))
	articles.Delete("/:id", handlers.RemoveArticle(service))
}
