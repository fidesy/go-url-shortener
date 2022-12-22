package models

import "time"

type URL struct {
	Hash           string    `pg:"hash"`
	OriginalURL    string    `pg:"original_url"`
	CreationDate   time.Time `pg:"creation_date"`
	ExpirationDate time.Time `pg:"expiration_date"`
}
