package repository

import (
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/fidesy/go-url-shortener/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	URL domain.URLRepository
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		URL: postgres.NewURLRepository(pool),
	}
}
