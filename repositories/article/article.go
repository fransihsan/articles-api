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

func (repo *ArticleRepository) GetAllArticles(author, keyword string) ([]A.Articles, error) {
	var articles []A.Articles
	if author != ""{
		repo.db.Order("created_at desc").Find(&articles, "author = ?", author)
		if len(articles) < 1 {
			return nil, errors.New("artikel berdasarkan nama pengarang tidak ditemukan")
		}
		return articles, nil
	}

	if keyword != "" {
		repo.db.Order("created_at desc").Find(&articles, "title LIKE ? OR body LIKE ?", "%" + keyword + "%", "%" + keyword + "%")
		if len(articles) < 1 {
			return nil, errors.New("kata kunci yang di cari tidak ditemukan")
		}
		return articles, nil
	}

	repo.db.Order("created_at desc").Find(&articles)
	if len(articles) < 1 {
		return nil, errors.New("tidak terdapat data artikel sama sekali")
	}
	return articles, nil
}