package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/fidesy/go-url-shortener/internal/repository"
	"github.com/fidesy/go-url-shortener/internal/repository/postgres"
	"github.com/fidesy/go-url-shortener/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	urls = []domain.URL{
		{OriginalURL: "https://google.com"},
		{OriginalURL: "https://amazon.com/"},
		{OriginalURL: "https://apple.com/some/path"},
	}
)

func GetRouter(t *testing.T) *gin.Engine {
	conf := config.DefaultConfig
	pool, err := postgres.NewPostgresPool(context.Background(), conf.Postgres)
	assert.Nil(t, err)

	repos := repository.NewRepository(pool)
	services := service.NewService(conf, repos)
	handler := NewHandler(services)

	return handler.InitRoutes()
}

func TestURLHandler_createShortURL(t *testing.T) {
	router := GetRouter(t)

	for i, url := range urls {
		body, _ := json.Marshal(domain.URL{
			OriginalURL: url.OriginalURL,
		})

		req, _ := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(body))

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		var responseBody struct {
			ShortURL string `json:"short_url"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		assert.Nil(t, err)

		urls[i].Hash = strings.Split(responseBody.ShortURL, "/")[3]
	}
}

func TestURLHandler_redirect(t *testing.T) {
	router := GetRouter(t)

	for _, url := range urls {
		req, _ := http.NewRequest(http.MethodGet, "/"+url.Hash, nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	}
}
