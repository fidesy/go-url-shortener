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
	api        = new(RestAPI)
	err  error = nil
	urls       = map[string]string{
		"https://google.com":                 "",
		"https://amazon.com":                 "",
		"https://ozon.ru":                    "",
		"https://yandex.ru/some/unique/path": "",
	}
)

func TestNewRestAPI(t *testing.T) {
	api, err = New(&RestAPIConfig{
		BindAddr: os.Getenv("BIND_ADDR"),
		DBURL:    os.Getenv("DBURL"),
		DBName:   os.Getenv("DBName"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, api)
}

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

func TestRedirect(t *testing.T) {
	w := httptest.NewRecorder()

	for _, short_url := range urls {
		hash := short_url[len(short_url)-7:]
		req, _ := http.NewRequest(http.MethodGet, "/"+hash, nil)

		api.router.ServeHTTP(w, req)
		assert.Equal(t, 301, w.Code)
		assert.Contains(t, w.Body.String(), "Moved Permanently")
	}

	api.db.Close()
}
