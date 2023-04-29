package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/fidesy/go-url-shortener/internal/repository"
	"github.com/fidesy/go-url-shortener/pkg/utils"
)

type URL interface {
	CreateShortURL(url domain.URL) string
	GetURLByHash(hash string) (domain.URL, error)
}

type URLService struct {
	conf *config.Config
	repo *repository.Repository
}

func NewURLService(conf *config.Config, repo *repository.Repository) *URLService {
	return &URLService{
		conf: conf,
		repo: repo,
	}
}

var _ URL = &URLService{}

func (s *URLService) CreateShortURL(url domain.URL) string {
	var (
		hash string
		err  error
	)

	url.CreationDate = time.Now().UTC()

	for {
		hash = utils.GenerateRandomSequence(6)
		url.Hash = hash
		_, err = s.repo.URL.CreateURL(context.Background(), url)
		if err != nil {
			log.Println(err)
			continue
		}

		break
	}

	if s.conf.Port != "" {
		s.conf.Port = ":" + s.conf.Port
	}

	shortURL := fmt.Sprintf("%s%s/%s", s.conf.Host, s.conf.Port, hash)

	return shortURL
}

func (s *URLService) GetURLByHash(hash string) (domain.URL, error) {
	return s.repo.URL.GetURLByHash(context.Background(), hash)
}
