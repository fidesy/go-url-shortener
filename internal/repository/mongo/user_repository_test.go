package mongo

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

func getUserRepository(t *testing.T) *UserRepository {
	cli, err := New(context.Background(), config.Default.Mongo)
	assert.Nil(t, err)

	repo := NewUserRepository(cli.Database("shortener").Collection("users"))

	return repo
}

func TestUserRepository_Create(t *testing.T) {
	repo := getUserRepository(t)

	id, err := repo.Create(context.Background(), user)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, id)
}

func TestUserRepository_GetByUsername(t *testing.T) {
	repo := getUserRepository(t)

	u, err := repo.Get(context.Background(), user.Username, user.Password)
	assert.Nil(t, err)
	assert.NotNil(t, u)

}

func TestUserRepository_UsernameExists(t *testing.T) {
	repo := getUserRepository(t)

	exists, err := repo.UsernameExists(context.Background(), user.Username)
	assert.Nil(t, err)
	assert.Equal(t, true, exists)
}

func TestUserRepository_Update(t *testing.T) {
	repo := getUserRepository(t)

	user.Name = updatedName
	err := repo.Update(context.Background(), user)
	assert.Nil(t, err)

	u, err := repo.Get(context.Background(), user.Username, user.Password)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, user.Name, u.Name)
}

func TestUserRepository_DeleteByUsername(t *testing.T) {
	repo := getUserRepository(t)

	err := repo.Delete(context.Background(), user.Username, user.Password)
	assert.Nil(t, err)

	_, err = repo.Get(context.Background(), user.Username, user.Password)
	assert.NotNil(t, err)
}
