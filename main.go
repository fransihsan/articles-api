package main

import (
	"articles-api/configs"
	_ArticleController "articles-api/deliveries/controllers/article"
	"articles-api/deliveries/routes"
	_ArticleRepo "articles-api/repositories/article"
	"articles-api/utils"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig(false)
	db := utils.InitDB(config)

	articleRepo := _ArticleRepo.NewArticleRepository(db)

	ac := _ArticleController.NewArticleController(articleRepo)

	e := echo.New()

	routes.RegisterPaths(e, ac)

	log.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
