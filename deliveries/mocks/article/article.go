package article

import (
	A "articles-api/entities/article"
	"errors"

	"gorm.io/gorm"
)

type MockArticleRepository struct{}

func (repo *MockArticleRepository) Create(newArticle A.Articles) (A.Articles, error) {
	return A.Articles{Model: gorm.Model{ID: 1}, Author: "frans", Title: "title", Body: "body"}, nil
}

func (repo *MockArticleRepository) GetAllArticles(author, query string) ([]A.Articles, error) {
	article1 := A.Articles{
		Model:  gorm.Model{ID: 1},
		Author: "frans",
		Title:  "title",
		Body:   "body",
	}

	article2 := A.Articles{
		Model:  gorm.Model{ID: 2},
		Author: "ihsan",
		Title:  "title2",
		Body:   "body2",
	}

	return []A.Articles{article2, article1}, nil
}

type MockFalseArticleRepository struct{}

func (repo *MockFalseArticleRepository) Create(newArticle A.Articles) (A.Articles, error) {
	return A.Articles{}, errors.New("gagal membuat artikel baru")
}

func (repo *MockFalseArticleRepository) GetAllArticles(author, query string) ([]A.Articles, error) {
	return nil, errors.New("tidak terdapat data artikel sama sekali")
}
