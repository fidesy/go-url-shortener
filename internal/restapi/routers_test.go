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
		Host:   os.Getenv("HOST"),
		Port:   os.Getenv("PORT"),
		DBURL:  os.Getenv("DB_URL"),
		DBName: os.Getenv("DB_NAME"),
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
	for url := range urls {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/create?url="+url, nil)
		api.router.ServeHTTP(w, req)
		assert.Equal(t, 200, w.Code)
		urls[url] = w.Body.String()
	}
}

func TestRedirect(t *testing.T) {
	for originalURL, shortURL := range urls {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, shortURL, nil)
		api.router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusPermanentRedirect, w.Code)
		// response text should contain html <a> tag with an url to redirect
		assert.Contains(t, w.Body.String(), originalURL)
	}

	api.db.Close()
}
