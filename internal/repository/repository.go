package repository

import (
	"context"
	"errors"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/fidesy/go-url-shortener/internal/repository/mongo"
	"github.com/fidesy/go-url-shortener/internal/repository/postgres"
)

type Repository struct {
	URL  domain.URLRepository
	User domain.UserRepository
}

func NewRepository(ctx context.Context, conf *config.Config) (*Repository, error) {
	var repos = new(Repository)

	switch conf.Database {
	case "postgres":
		pool, err := postgres.NewPool(ctx, conf.Postgres)
		if err != nil {
			return nil, err
		}

		// Come up with another way to implement this
		_repos := postgres.NewRepository(pool)
		repos.URL = _repos.URL
		repos.User = _repos.User

	case "mongo":
		mongocli, err := mongo.New(ctx, conf.Mongo)
		if err != nil {
			return nil, err
		}

		_repos := mongo.NewRepository(mongocli)
		repos.URL = _repos.URL
		repos.User = _repos.User

	default:
		return nil, errors.New("unknown config.database field; options: postgres, mongo")
	}

	return repos, nil
}
