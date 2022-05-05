package gocache

import (
	A "articles-api/entities/article"
	"time"

	"github.com/patrickmn/go-cache"
)
var Cache = cache.New(1*time.Minute, 5*time.Minute)
type Emp []struct {
    ID        uint   `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func SetCache(key string, articles []A.Articles) bool {
	responses := make(Emp, len(articles))
	for i := 0; i < len(articles); i++ {
		responses[i].ID = articles[i].ID
		responses[i].Author = articles[i].Author
		responses[i].Title = articles[i].Title
		responses[i].Body = articles[i].Body
		responses[i].CreatedAt = articles[i].CreatedAt.Format("02 January 2006 15:04")
	}
    Cache.Set(key, responses, cache.DefaultExpiration)
    return true
}
	
func GetCache(key string) (Emp, bool) {
    var emp Emp
    var found bool
    data, found := Cache.Get(key)
    if found {
      emp = data.(Emp)
    }
    return emp, found
}