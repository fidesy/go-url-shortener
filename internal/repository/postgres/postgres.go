package postgres

import (
	"context"
	"fmt"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(ctx context.Context, conf config.Postgres) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			conf.Username,
			conf.Password,
			conf.Host,
			conf.Port,
			conf.DBName,
			conf.SSLMode,
		))
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}

type Repository struct {
	URL  domain.URLRepository
	User domain.UserRepository
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		URL:  NewURLRepository(pool),
		User: NewUserRepository(pool),
	}
}
