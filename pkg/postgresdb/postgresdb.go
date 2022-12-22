package postgresdb

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/fidesy/go-url-shortener/pkg/shortener"
)

type PostgresDB struct {
	db  *sqlx.DB
}

func New() *PostgresDB {
	return &PostgresDB{}
}

const initScheme = `
CREATE TABLE IF NOT EXISTS urls (
    hash VARCHAR(6) PRIMARY KEY,
    original_url VARCHAR,
    creation_date DATE,
    expiration_date DATE
);`

func (p *PostgresDB) Open(ctx context.Context, dbURI string) error {
	db, err := sqlx.Open("postgres", dbURI)
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, initScheme)
	if err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *PostgresDB) Close() {
	p.db.Close()
}

const (
	insertTemplate = "INSERT INTO urls VALUES($1, $2, $3, $4)"
	selectTemplate = "SELECT original_url FROM urls WHERE hash=$1"
)

func (p *PostgresDB) CreateURL(ctx context.Context, originalURL string) (string, error) {
	var hash string
	for {
		hash = shortener.GetRandomSequence(6)

		// check if this hash already exists
		if _, err := p.GetOriginalURL(ctx, hash); err != nil {
			break
		}
	}

	_, err := p.db.ExecContext(ctx, insertTemplate, hash, originalURL, time.Now(), nil)
	return hash, err
}

func (p *PostgresDB) GetOriginalURL(ctx context.Context, hash string) (string, error) {
	var original_url []string
	err := p.db.SelectContext(ctx, &original_url, selectTemplate, hash)
	if len(original_url) == 0 {
		return "", errors.New("not found")
	}

	return original_url[0], err
}
