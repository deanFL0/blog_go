package routes

import (
	"github.com/deanFL0/blog_api_go/api/handlers"
	"github.com/deanFL0/blog_api_go/pkg/article"
	"github.com/gofiber/fiber/v2"
)

func ArticleRouter(app fiber.Router, service article.Service) {
	app.Get("/articles", handlers.GetArticles(service))
	app.Post("/articles", handlers.AddArticle(service))
	app.Put("/articles", handlers.UpdateArticle(service))
	app.Delete("/articles", handlers.RemoveArticle(service))
}
