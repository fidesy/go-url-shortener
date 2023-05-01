package mongo

import (
	"context"
	"fmt"

	"github.com/fidesy/go-url-shortener/internal/config"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(ctx context.Context, conf config.Mongo) (*mongo.Client, error) {
	if conf.Port != "" {
		conf.Host += ":" + conf.Port
	}

	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s",
		conf.Username,
		conf.Password,
		conf.Host,
	))

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	return client, nil
}

type Repository struct {
	URL  domain.URLRepository
	User domain.UserRepository
}

func NewRepository(cli *mongo.Client) *Repository {
	return &Repository{
		URL:  NewURLRepository(cli.Database("shortener").Collection("urls")),
		User: NewUserRepository(cli.Database("shortener").Collection("users")),
	}
}
