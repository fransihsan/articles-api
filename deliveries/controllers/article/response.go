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

type ResponseGetAllArticles struct {
	ID        uint      `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func ToResponseGetAllArticles(Articles []A.Articles) []ResponseGetAllArticles {
	responses := make([]ResponseGetAllArticles, len(Articles))
	for i := 0; i < len(Articles); i++ {
		responses[i].ID = Articles[i].ID
		responses[i].Author = Articles[i].Author
		responses[i].Title = Articles[i].Title
		responses[i].Body = Articles[i].Body
		responses[i].CreatedAt = Articles[i].CreatedAt
	}
	return responses
}
