package domain

import (
	"context"
	"time"
)

type URL struct {
	ID             int       `json:"id" db:"id"`
	Hash           string    `json:"hash" db:"hash"`
	OriginalURL    string    `json:"original_url" db:"original_url" binding:"required"`
	CreationDate   time.Time `json:"creation_date" db:"creation_date"`
	ExpirationDate time.Time `json:"expiration_date" db:"expiration_date"`
}

type URLRepository interface {
	CreateURL(ctx context.Context, url URL) (int, error)
	GetURLByHash(ctx context.Context, hash string) (URL, error)
}
