package postgres

import (
	"context"

	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository struct {
	pool *pgxpool.Pool
}

func NewURLRepository(pool *pgxpool.Pool) *URLRepository {
	return &URLRepository{pool: pool}
}

var _ domain.URLRepository = &URLRepository{}

func (r *URLRepository) CreateURL(ctx context.Context, url domain.URL) (interface{}, error) {
	var id int
	err := r.pool.QueryRow(
		ctx,
		"INSERT INTO urls(user_id, hash, original_url, creation_date, expiration_date) VALUES($1, $2, $3, $4, $5) RETURNING id",
		url.UserID,
		url.Hash,
		url.OriginalURL,
		url.CreationDate,
		url.ExpirationDate,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *URLRepository) GetURLByHash(ctx context.Context, hash string) (domain.URL, error) {
	var url domain.URL
	err := r.pool.QueryRow(
		ctx,
		"SELECT id, hash, original_url, creation_date, expiration_date FROM urls WHERE hash=$1",
		hash,
	).Scan(
		&url.ID,
		&url.Hash,
		&url.OriginalURL,
		&url.CreationDate,
		&url.ExpirationDate,
	)
	if err != nil {
		return domain.URL{}, err
	}

	return url, nil
}
