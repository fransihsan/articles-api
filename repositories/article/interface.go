package article

import A "articles-api/entities/article"

type Article interface {
	Create(newArticle A.Articles) (A.Articles, error)
	GetAllArticles() ([]A.Articles, error)
	FilteByAuthorName(author string) ([]A.Articles, error)
	SearchArticle(query string) ([]A.Articles, error)
}