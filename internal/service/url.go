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
	conf config.Config
	repo *repository.Repository
}

func NewURLService(conf config.Config, repo *repository.Repository) *URLService {
	s := &URLService{
		conf: conf,
		repo: repo,
	}

	if s.conf.Port != "" {
		s.conf.Port = ":" + s.conf.Port
	}

	return s
}

var _ URL = &URLService{}

func (s *URLService) CreateShortURL(url domain.URL) string {
	var (
		hash string
		err  error
	)

	for {
		hash = utils.GenerateRandomSequence(6)
		_, err = s.repo.URL.GetURLByHash(context.Background(), hash)
		// hash already exists
		if err == nil {
			continue
		}

		url.Hash = hash
		url.CreationDate = time.Now().UTC()

		_, err = s.repo.URL.CreateURL(context.Background(), url)
		if err != nil {
			log.Println(err)
		}

		break
	}

	shortURL := fmt.Sprintf("%s%s/%s", s.conf.Host, s.conf.Port, hash)

	return shortURL
}

func (s *URLService) GetURLByHash(hash string) (domain.URL, error) {
	return s.repo.URL.GetURLByHash(context.Background(), hash)
}
