package database

import (
	"context"
	"errors"
	"time"

	"github.com/fidesy/go-url-shortener/pkg/shortener"
	"github.com/jackc/pgx/v5"
)

type PostgreSQL struct {
	conn *pgx.Conn
}

func NewPostgreSQL() *PostgreSQL {
	return &PostgreSQL{}
}

const initScheme = `
CREATE TABLE IF NOT EXISTS urls (
    hash            VARCHAR(6) PRIMARY KEY,
    original_url    VARCHAR,
    creation_date   DATE,
    expiration_date DATE
);`

func (p *PostgreSQL) Open(ctx context.Context, DBURL string) error {
	connection, err := pgx.Connect(ctx, DBURL)
	if err != nil {
		return err
	}

	if err := connection.Ping(ctx); err != nil {
		return err
	}

	if _, err = connection.Exec(ctx, initScheme); err != nil {
		return err
	}

	p.conn = connection

	return nil
}

func (p *PostgreSQL) Close(ctx context.Context) error {
	return p.conn.Close(ctx)
}

const (
	insertTemplate = "INSERT INTO urls VALUES($1, $2, $3, $4)"
	selectTemplate = "SELECT original_url FROM urls WHERE hash=$1"
)

func (p *PostgreSQL) CreateShortURL(ctx context.Context, originalURL string) (string, error) {
	var hash string
	for {
		hash = shortener.GetRandomSequence(6)

		// check if this hash already exists
		if _, err := p.GetOriginalURL(ctx, hash); err != nil {
			break
		}
	}

	_, err := p.conn.Exec(ctx, insertTemplate, hash, originalURL, time.Now(), nil)
	return hash, err
}

func (p *PostgreSQL) GetOriginalURL(ctx context.Context, hash string) (string, error) {
	rows, err := p.conn.Query(ctx, selectTemplate, hash)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var originalURLs []string
	for rows.Next() {
		var originalURL string
		if err = rows.Scan(&originalURL); err != nil {
			break
		}

		originalURLs = append(originalURLs, originalURL)
	}

	if len(originalURLs) == 0 {
		return "", errors.New("not found")
	}

	return originalURLs[0], err
}
