package article

import (
	"articles-api/deliveries/controllers/common"
	_ArticleRepo "articles-api/repositories/article"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ArticleController struct {
	repo _ArticleRepo.Article
}

func NewArticleController(repository _ArticleRepo.Article) *ArticleController {
	return &ArticleController{
		repo: repository,
	}
}

func (ctl *ArticleController) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newArticle RequestCreateArticle

		if err := c.Bind(&newArticle); err != nil || strings.TrimSpace(newArticle.Author) == "" || strings.TrimSpace(newArticle.Title) == "" || strings.TrimSpace(newArticle.Body) == "" {
			return c.JSON(http.StatusBadRequest, common.BadRequest("input dari author tidak sesuai, author, title atau body tidak boleh kosong"))
		}

		res, err := ctl.repo.Create(newArticle.ToEntityArticle())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "sukses menambahkan artikel baru", ToResponseCreateArticle(res)))
	}
}

func (ctl *ArticleController) GetAllArticles() echo.HandlerFunc {
	return func(c echo.Context) error {
		keyword := c.QueryParam("query")
		author := c.QueryParam("author")
		res, err := ctl.repo.GetAllArticles(author, keyword)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError(err.Error()))
		}
		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "sukses mendapatkan semua artikel", ToResponseGetAllArticles(res)))
	}
}