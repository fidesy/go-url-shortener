package database

import (
	"context"
	"errors"
)

type Database interface {
	Open(context.Context, string) error
	Close()
	CreateShortURL(context.Context, string) (string, error)
	GetOriginalURL(context.Context, string) (string, error)
}

func NewDatabase(DBName string) (Database, error) {
	switch DBName {
	case "postgresql":
		return &PostgreSQL{}, nil
	default:
		return nil, errors.New("unknown name of database, options: 'postgresql'")
	}
}
