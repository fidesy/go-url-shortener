package domain

import "context"

type User struct {
	ID       interface{} `json:"-" bson:"_id,omitempty" db:"id"`
	Name     string      `json:"name" bson:"name" db:"name" binding:"required"`
	Username string      `json:"username" bson:"username" db:"username" binding:"required"`
	Password string      `json:"password" bson:"password_hash" db:"password_hash" binding:"required"`
}

type UserRepository interface {
	Create(ctx context.Context, user User) (interface{}, error)
	Get(ctx context.Context, username, password string) (User, error)
	Update(ctx context.Context, user User) error
	Delete(ctx context.Context, username, password string) error
	UsernameExists(ctx context.Context, username string) (bool, error)
}
