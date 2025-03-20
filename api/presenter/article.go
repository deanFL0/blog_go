package presenter

import (
	"github.com/deanFL0/blog_api_go/pkg/entities"
	"github.com/gofiber/fiber/v2"
)

type Article struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"user_id"`
}

func ArticleSuccessResponse(data *entities.Article) *fiber.Map {
	article := Article{
		ID:    data.ID,
		Title: data.Title,
		Body:  data.Body,
	}
	return &fiber.Map{
		"status": true,
		"data":   article,
		"error":  nil,
	}
}

func ArticlesSuccessResponse(data *[]entities.Article) *fiber.Map {
	var articles []Article
	for _, article := range *data {
		articles = append(articles, Article{
			ID:    article.ID,
			Title: article.Title,
			Body:  article.Body,
		})
	}
	return &fiber.Map{
		"status": true,
		"data":   articles,
		"error":  nil,
	}
}

func ArticleErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   "",
		"error":  err.Error(),
	}
}
