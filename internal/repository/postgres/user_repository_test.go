package postgres

import (
	"context"
	"testing"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	user = domain.User{
		Name:     "John",
		Username: "johndoe",
		Password: "mysecretpass",
	}
	updatedName = "Dave"
)

func GetUserRepository(t *testing.T) *UserRepository {
	pool, err := NewPostgresPool(context.Background(), config.DefaultConfig.Postgres)
	assert.Nil(t, err)

	repo := NewUserRepository(pool)

	return repo
}

func TestUserRepository_Create(t *testing.T) {
	repo := GetUserRepository(t)

	id, err := repo.Create(context.Background(), user)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}

func TestUserRepository_GetByUsername(t *testing.T) {
	repo := GetUserRepository(t)

	u, err := repo.Get(context.Background(), user.Username, user.Password)
	assert.Nil(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_UsernameExists(t *testing.T) {
	repo := GetUserRepository(t)

	exists, err := repo.UsernameExists(context.Background(), user.Username)
	assert.Nil(t, err)
	assert.Equal(t, true, exists)
}

func TestUserRepository_Update(t *testing.T) {
	repo := GetUserRepository(t)

	user.Name = updatedName
	err := repo.Update(context.Background(), user)
	assert.Nil(t, err)

	u, err := repo.Get(context.Background(), user.Username, user.Password)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, user.Name, u.Name)
}

func TestUserRepository_DeleteByUsername(t *testing.T) {
	repo := GetUserRepository(t)

	err := repo.Delete(context.Background(), user.Username, user.Password)
	assert.Nil(t, err)

	_, err = repo.Get(context.Background(), user.Username, user.Password)
	assert.NotNil(t, err)
}

