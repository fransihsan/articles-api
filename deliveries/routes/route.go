package routes

import (
	"articles-api/deliveries/controllers/article"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterPaths(e *echo.Echo, ac *article.ArticleController) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// User Route
	e.POST("/articles", ac.Create())
	e.GET("/articles", ac.GetAllArticles())
}
