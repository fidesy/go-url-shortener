package restapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	api = New(&RestAPIConfig{
		BindAddr: os.Getenv("BIND_ADDR"),
		DBURL:    os.Getenv("DBURL"),
	})
	urls = map[string]string{
		"https://google.com":                 "",
		"https://amazon.com":                 "",
		"https://ozon.ru":                    "",
		"https://yandex.ru/some/unique/path": "",
	}
)

func TestConfigureRoutersAndDatabase(t *testing.T) {
	api.configureRouters()
	err := api.configureDatabase(context.Background())
	assert.Nil(t, err)
}

func TestCreateShortURL(t *testing.T) {
	w := httptest.NewRecorder()

	for url := range urls {
		req, _ := http.NewRequest(http.MethodPost, "/create?url="+url, nil)
		api.router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		urls[url] = w.Body.String()	
	}
}

// func TestCloseDatabase(t *testing.T) {
// 	err := api.db.Close(context.Background())
// 	assert.Nil(t, err)
// }
