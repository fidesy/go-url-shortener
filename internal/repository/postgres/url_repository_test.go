package postgres

import (
	"context"
	"testing"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/fidesy/go-url-shortener/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var (
	urls = []domain.URL{
		{OriginalURL: "https://google.com", Hash: utils.GenerateRandomSequence(6)},
		{OriginalURL: "https://amazon.com/", Hash: utils.GenerateRandomSequence(6)},
		{OriginalURL: "https://apple.com/some/path", Hash: utils.GenerateRandomSequence(6)},
	}
)

func GetURLRepository(t *testing.T) *URLRepository {
	pool, err := NewPostgresPool(context.Background(), config.DefaultConfig.Postgres)
	assert.Nil(t, err)

	repo := NewURLRepository(pool)

	return repo
}

func TestURLRepository_CreateURL(t *testing.T) {
	repo := GetURLRepository(t)

	for _, url := range urls {
		id, err := repo.CreateURL(
			context.Background(),
			domain.URL{
				OriginalURL: url.OriginalURL,
				Hash:        url.Hash,
			})

		assert.Nil(t, err)
		assert.NotEqual(t, 0, id)
	}
}

func TestURLRepository_GetURLByHash(t *testing.T) {
	repo := GetURLRepository(t)

	for _, url := range urls {
		_url, err := repo.GetURLByHash(context.Background(), url.Hash)
		assert.Nil(t, err)
		assert.Equal(t, url.OriginalURL, _url.OriginalURL)
	}
}
