package article

import (
	A "articles-api/entities/article"
	"time"
)

type ResponseCreateArticle struct {
	ID        uint      `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func ToResponseCreateArticle(Article A.Articles) ResponseCreateArticle {
	return ResponseCreateArticle{
		ID:        Article.ID,
		Author:    Article.Author,
		Title:     Article.Title,
		Body:      Article.Body,
		CreatedAt: Article.CreatedAt,
	}
}
