package article

import (
	"bytes"
	"encoding/json"
	"articles-api/deliveries/controllers/common"
	MockArticle "articles-api/deliveries/mocks/article"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/assert"
)

var (
	e        = echo.New()
	rootPath = "/articles"
)

func TestCreate(t *testing.T) {
	t.Run("fail to bind json", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateArticle{
			Author: "",
			Title: "",
			Body: "",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		articleController := NewArticleController(&MockArticle.MockArticleRepository{})
		articleController.Create()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("fail to create", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateArticle{
			Author: "frans",
			Title: "title",
			Body: "body",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)

		articleController := NewArticleController(&MockArticle.MockFalseArticleRepository{})
		articleController.Create()(context)
		context.SetPath(rootPath)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Nil(t, response.Data)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to create", func(t *testing.T) {
		requestBody, _ := json.Marshal(RequestCreateArticle{
			Author: "frans",
			Title: "title",
			Body: "body",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath(rootPath)

		articleController := NewArticleController(&MockArticle.MockArticleRepository{})
		articleController.Create()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.NotNil(t, response.Data)
		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetAllUsers(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		log.Info("tidak dapat memuat env file", err)
	}

	t.Run("fail to get all articles", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		articleController := NewArticleController(&MockArticle.MockFalseArticleRepository{})
		articleController.GetAllArticles()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})

	t.Run("succeed to get all users", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(nil))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath(fmt.Sprintf("%v", rootPath))

		articleController := NewArticleController(&MockArticle.MockArticleRepository{})
		articleController.GetAllArticles()(context)

		response := common.Response{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}