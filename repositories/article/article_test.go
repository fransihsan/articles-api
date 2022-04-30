package article

import (
	"articles-api/configs"
	A "articles-api/entities/article"
	ArticleMock "articles-api/repositories/mock/article"
	"articles-api/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	config = configs.GetConfig(true)
	db     = utils.InitDB(config)
)

func Migrator() {
	db.Migrator().DropTable(&A.Articles{})

	db.AutoMigrate(&A.Articles{})
}

func TestCreate(t *testing.T) {
	Migrator()
	repo := NewArticleRepository(db)

	mockArticle := ArticleMock.MockArticle()

	t.Run("positive", func(t *testing.T) {
		res, err := repo.Create(mockArticle)
		assert.Nil(t, err)
		assert.Equal(t, mockArticle.Author, res.Author)
	})

	t.Run("negative", func(t *testing.T) {
		mockArticle.ID = 1
		_, err := repo.Create(mockArticle)
		assert.NotNil(t, err)
	})
}