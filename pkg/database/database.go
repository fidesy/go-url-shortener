package database

import "context"

type Database interface {
	Open(context.Context, string) error
	Close()
	CreateShortURL(context.Context, string) (string, error)
	GetOriginalURL(context.Context, string) (string, error)
}
