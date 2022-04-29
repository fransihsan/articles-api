package article

import A "articles-api/entities/article"

type Article interface {
	Create(newArticle A.Articles) (A.Articles, error)
	GetAllArticles(author, keyword string) ([]A.Articles, error)
}