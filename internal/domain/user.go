package domain

import "context"

type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name" binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

type UserRepository interface {
	Create(ctx context.Context, user User) (int, error)
	Get(ctx context.Context, username, password string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, username, password string) error
	UsernameExists(ctx context.Context, username string) (bool, error)
}
