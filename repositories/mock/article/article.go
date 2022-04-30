package article

import (
	A "articles-api/entities/article"
)

func MockArticle() A.Articles {
	mockArticle := A.Articles{
		Author: "frans",
		Title:  "title",
		Body:   "body",
	}
	return mockArticle
}
