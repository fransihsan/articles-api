package article

import (
	"errors"
	A "articles-api/entities/article"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		db: db,
	}
}

func (repo *ArticleRepository) Create(newArticle A.Articles) (A.Articles, error) {
	if err := repo.db.Create(&newArticle).Error; err != nil {
		log.Warn(err)
		return A.Articles{}, errors.New("gagal membuat artikel baru")
	}
	return newArticle, nil
}