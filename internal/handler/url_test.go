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
	urlUser = domain.User{
		Name:     "User",
		Username: "urluser",
		Password: "urluser",
	}
)

func getRouter(t *testing.T) *gin.Engine {
	conf := config.Default

	repos, err := repository.NewRepository(context.Background(), config.Default)
	assert.Nil(t, err)

	services := service.NewService(conf, repos)
	handler := NewHandler(services)

	return handler.InitRoutes()
}

func getAuthorizationToken(t *testing.T) string {
	router := getRouter(t)

	body, _ := json.Marshal(urlUser)
	req, _ := http.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Sign In
	req, _ = http.NewRequest(http.MethodPost, "/auth/sign-in", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	type responseBody struct {
		Token string `json:"token"`
	}

	var respBody responseBody
	err := json.Unmarshal(w.Body.Bytes(), &respBody)
	assert.Nil(t, err)

	return respBody.Token
}

func TestURLHandler_createShortURL(t *testing.T) {
	router := getRouter(t)
	token := getAuthorizationToken(t)

	for i, url := range urls {
		body, _ := json.Marshal(domain.URL{
			OriginalURL: url.OriginalURL,
		})

		req, _ := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(body))
		req.Header.Add("Authorization", "Bearer "+token)
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
	router := getRouter(t)

	for _, url := range urls {
		req, _ := http.NewRequest(http.MethodGet, "/"+url.Hash, nil)

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
	}
}
