package service

import (
	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/repository"
)

type Service struct {
	URL
}

func NewService(conf *config.Config, repos *repository.Repository) *Service {
	return &Service{
		URL: NewURLService(conf, repos),
	}
}
