package postgresql

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	db   = New()
	urls = map[string]string{
		"https://google.com":                 "",
		"https://amazon.com":                 "",
		"https://ozon.ru":                    "",
		"https://yandex.ru/some/unique/path": "",
	}
)

func TestOpenPostgresDB(t *testing.T) {
	err := db.Open(context.Background(), os.Getenv("DB_URL"))
	assert.Nil(t, err)
}

func TestCreateURL(t *testing.T) {
	for url := range urls {
		hash, err := db.CreateShortURL(context.Background(), url)
		assert.Nil(t, err)
		urls[url] = hash
	}
}

func TestGetOriginalURL(t *testing.T) {
	for originalURL, hash := range urls {
		url, err := db.GetOriginalURL(context.Background(), hash)
		assert.Nil(t, err)
		assert.Equal(t, originalURL, url)
	}
}

func TestCloseDatabase(t *testing.T) {
	db.Close()
}
