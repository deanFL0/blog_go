package article

import (
	"github.com/deanFL0/blog_api_go/pkg/entities"
)

type Service interface {
	InsertArticle(article *entities.Article) (*entities.Article, error)
	FetchArticle(ID int) (*entities.Article, error)
	FetchArticles() (*[]entities.Article, error)
	UpdateArticle(ID int, article *entities.Article) (*entities.Article, error)
	RemoveArticle(ID int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) InsertArticle(article *entities.Article) (*entities.Article, error) {
	return s.repository.CreteArticle(article)
}

func (s *service) FetchArticle(ID int) (*entities.Article, error) {
	return s.repository.ReadArticle(ID)
}

func (s *service) FetchArticles() (*[]entities.Article, error) {
	return s.repository.ReadArticles()
}

func (s *service) UpdateArticle(ID int, article *entities.Article) (*entities.Article, error) {
	return s.repository.UpdateArticle(ID, article)
}

func (s *service) RemoveArticle(ID int) error {
	return s.repository.DeleteArticle(ID)
}
