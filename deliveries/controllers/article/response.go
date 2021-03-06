package article

import (
	A "articles-api/entities/article"
)

type ResponseCreateArticle struct {
	ID        uint   `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func ToResponseCreateArticle(Article A.Articles) ResponseCreateArticle {
	return ResponseCreateArticle{
		ID:        Article.ID,
		Author:    Article.Author,
		Title:     Article.Title,
		Body:      Article.Body,
		CreatedAt: Article.CreatedAt.Format("02 January 2006 15:04"),
	}
}

type ResponseGetAllArticles struct {
	ID        uint   `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func ToResponseGetAllArticles(Articles []A.Articles) []ResponseGetAllArticles {
	responses := make([]ResponseGetAllArticles, len(Articles))
	for i := 0; i < len(Articles); i++ {
		responses[i].ID = Articles[i].ID
		responses[i].Author = Articles[i].Author
		responses[i].Title = Articles[i].Title
		responses[i].Body = Articles[i].Body
		responses[i].CreatedAt = Articles[i].CreatedAt.Format("02 January 2006 15:04")
	}
	return responses
}
