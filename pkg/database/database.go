package database

import "context"

type Database interface {
	Open(context.Context, string) error
	Close(context.Context) error
	CreateShortURL(context.Context, string) (string, error)
	GetOriginalURL(context.Context, string) (string, error)
}
