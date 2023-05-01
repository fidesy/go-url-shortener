package postgres

import (
	"context"
	"errors"

	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

var _ domain.UserRepository = &UserRepository{}

func (r *UserRepository) Create(ctx context.Context, user domain.User) (interface{}, error) {
	var id int
	err := r.pool.QueryRow(
		ctx,
		"INSERT INTO users(name, username, password_hash) VALUES($1, $2, $3) RETURNING id",
		user.Name,
		user.Username,
		user.Password,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) Get(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User
	err := r.pool.QueryRow(
		ctx,
		"SELECT id, name, username, password_hash FROM users WHERE username=$1 AND password_hash=$2",
		username,
		password,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
	)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) error {
	_, err := r.pool.Exec(
		ctx,
		"UPDATE users SET name=$1, username=$2, password_hash=$3 WHERE username=$2 AND password_hash=$3",
		user.Name,
		user.Username,
		user.Password,
	)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, username, password string) error {
	_, err := r.pool.Exec(
		ctx,
		"DELETE FROM users WHERE username=$1 AND password_hash=$2",
		username,
		password,
	)
	return err
}

func (r *UserRepository) UsernameExists(ctx context.Context, username string) (bool, error) {
	var id int
	err := r.pool.QueryRow(
		ctx,
		"SELECT id FROM users WHERE username=$1",
		username,
	).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return !(id == 0), nil
}
