package handler

import (
	"github.com/fidesy/go-url-shortener/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/:hash", h.redirect)
	router.POST("/create", h.createShortURL)

	return router
}