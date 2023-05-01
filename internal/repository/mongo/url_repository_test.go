package mongo

import (
	"context"
	"testing"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/fidesy/go-url-shortener/pkg/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	urls = []domain.URL{
		{UserID: primitive.NewObjectID(), OriginalURL: "https://google.com", Hash: utils.GenerateRandomSequence(6)},
		{UserID: primitive.NewObjectID(), OriginalURL: "https://amazon.com/", Hash: utils.GenerateRandomSequence(6)},
		{UserID: primitive.NewObjectID(), OriginalURL: "https://apple.com/some/path", Hash: utils.GenerateRandomSequence(6)},
	}
)

func getURLRepository(t *testing.T) *URLRepository {
	cli, err := New(context.Background(), config.Default.Mongo)
	assert.Nil(t, err)

	repo := NewURLRepository(cli.Database("shortener").Collection("urls"))

	return repo
}
func TestMongoRepository_CreateURL(t *testing.T) {
	repo := getURLRepository(t)

	for _, url := range urls {
		id, err := repo.CreateURL(
			context.Background(),
			url,
		)
		assert.Nil(t, err)
		assert.NotEqual(t, 0, id)
	}
}

func TestMongoRepository_GetURLByHash(t *testing.T) {
	repo := getURLRepository(t)

	for _, url := range urls {
		_url, err := repo.GetURLByHash(context.Background(), url.Hash)
		assert.Nil(t, err)
		assert.Equal(t, url.OriginalURL, _url.OriginalURL)
	}
}
