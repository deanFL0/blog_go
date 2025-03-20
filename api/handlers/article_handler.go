package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/deanFL0/blog_api_go/api/presenter"
	"github.com/deanFL0/blog_api_go/pkg/article"
	"github.com/deanFL0/blog_api_go/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

func AddArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody entities.Article
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		if requestBody.Title == "" || requestBody.Body == "" {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(errors.New("please specify title and body")))
		}
		result, err := service.InsertArticle(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticleSuccessResponse(result))
	}
}

func UpdateArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}

		existingArticle, err := service.FetchArticle(int(ID))
		if err != nil {
			return c.JSON(presenter.ArticleErrorResponse(err))
		}

		var requestBody entities.Article
		err = c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}

		if requestBody.Title != "" {
			existingArticle.Title = requestBody.Title
		}
		if requestBody.Body != "" {
			existingArticle.Body = requestBody.Body
		}

		result, err := service.UpdateArticle(ID, existingArticle)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticleSuccessResponse(result))
	}
}

func RemoveArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		err = service.RemoveArticle(int(ID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(&fiber.Map{
			"status": true,
			"data":   "Deleted successfully",
			"err":    nil,
		})
	}
}

func GetArticle(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		fetched, err := service.FetchArticle(int(ID))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticleSuccessResponse(fetched))
	}
}

func GetArticles(service article.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fetched, err := service.FetchArticles()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(presenter.ArticleErrorResponse(err))
		}
		return c.JSON(presenter.ArticlesSuccessResponse(fetched))
	}
}
