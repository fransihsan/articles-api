package article

import (
	A "articles-api/entities/article"
)

type RequestCreateArticle struct {
	Author string `json:"author"`
	Title  string `json:"title" `
	Body   string `json:"body"`
}

func (Req RequestCreateArticle) ToEntityArticle() A.Articles {
	return A.Articles{
		Author: Req.Author,
		Title:  Req.Title,
		Body:   Req.Body,
	}
}
