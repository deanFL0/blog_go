package article

import (
	"github.com/deanFL0/blog_api_go/pkg/entities"
	"gorm.io/gorm"
)

type Repository interface {
	CreteArticle(article *entities.Article) (*entities.Article, error)
	ReadArticle(ID int) (*entities.Article, error)
	ReadArticles() (*[]entities.Article, error)
	UpdateArticle(ID int, article *entities.Article) (*entities.Article, error)
	DeleteArticle(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreteArticle(article *entities.Article) (*entities.Article, error) {
	return article, r.db.Create(article).Error
}

func (r *repository) ReadArticle(ID int) (*entities.Article, error) {
	var article entities.Article
	result := r.db.Find(&article, ID)
	return &article, result.Error
}

func (r *repository) ReadArticles() (*[]entities.Article, error) {
	var article []entities.Article
	result := r.db.Find(&article)
	return &article, result.Error
}

func (r *repository) UpdateArticle(ID int, article *entities.Article) (*entities.Article, error) {
	existingArticle, err := r.ReadArticle(article.ID)
	if err != nil {
		return nil, err
	}

	existingArticle.Title = article.Title
	existingArticle.Body = article.Body

	return existingArticle, r.db.Save(existingArticle).Error
}

func (r *repository) DeleteArticle(ID int) error {
	var article entities.Article
	err := r.db.Find(&article, ID).Error
	if err != nil {
		return err
	}
	return r.db.Delete(&article).Error
}
