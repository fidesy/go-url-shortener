package database

import (
	"context"
	"errors"
	"github.com/fidesy/go-url-shortener/internal/databases/postgresql"
)

type Database interface {
	Open(context.Context, string) error
	Close()
	CreateShortURL(context.Context, string) (string, error)
	GetOriginalURL(context.Context, string) (string, error)
}

func New(DBName string) (Database, error) {
	switch DBName {
	case "postgresql":
		return postgresql.New(), nil
	default:
		return nil, errors.New("unknown name of database, options: 'postgresql'")
	}
}
