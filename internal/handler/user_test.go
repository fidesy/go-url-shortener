package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	user = domain.User{
		Name:     "John",
		Username: "johndoe",
		Password: "johndoe",
	}
)

func TestUserHandler_signUp(t *testing.T) {
	router := GetRouter(t)
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/auth/sign-up", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUserHandler_signIn(t *testing.T) {
	router := GetRouter(t)
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/auth/sign-in", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
