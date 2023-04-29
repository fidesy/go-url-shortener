package handler

import (
	"net/http"

	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) redirect(c *gin.Context) {
	hash := c.Param("hash")
	url, err := h.services.URL.GetURLByHash(hash)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url.OriginalURL)
}

func (h *Handler) createShortURL(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input domain.URL

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	input.UserID = userID

	shortURL := h.services.URL.CreateShortURL(input)

	c.JSON(http.StatusCreated, map[string]interface{}{
		"short_url": shortURL,
	})
}
