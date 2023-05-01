package domain

import (
	"context"
	"time"
)

type URL struct {
	ID             interface{}       `json:"-" bson:"_id,omitempty" db:"id"`
	UserID         interface{}       `json:"-" bson:"user_id" db:"user_id"`
	Hash           string    `json:"hash" bson:"hash" db:"hash"`
	OriginalURL    string    `json:"original_url" bson:"original_url" db:"original_url" binding:"required"`
	CreationDate   time.Time `json:"creation_date" bson:"creation_date" db:"creation_date"`
	ExpirationDate time.Time `json:"expiration_date" bson:"expiration_date" db:"expiration_date"`
}

type URLRepository interface {
	CreateURL(ctx context.Context, url URL) (interface{}, error)
	GetURLByHash(ctx context.Context, hash string) (URL, error)
}
