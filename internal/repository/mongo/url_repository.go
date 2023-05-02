package mongo

import (
	"context"
	"errors"

	"github.com/fidesy/go-url-shortener/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLRepository struct {
	coll *mongo.Collection
}

func NewURLRepository(coll *mongo.Collection) *URLRepository {
	return &URLRepository{coll: coll}
}

var _ domain.URLRepository = &URLRepository{}

func (r *URLRepository) CreateURL(ctx context.Context, url domain.URL) (interface{}, error) {
	// we should place this code here only because we have the same model struct
	// for 2 different databases
	_, ok := url.UserID.(primitive.ObjectID)
	if !ok {
		userIDInt, ok := url.UserID.(string)
		if !ok {
			return nil, errors.New("url user_id has an invalid type")
		}

		userID, err := primitive.ObjectIDFromHex(userIDInt)
		if err != nil {
			return nil, err
		}
		url.UserID = userID
	}
	// end of shitty code

	res, err := r.coll.InsertOne(ctx, url)
	if err != nil {
		return 0, err
	}

	return res.InsertedID, nil
}

func (r *URLRepository) GetURLByHash(ctx context.Context, hash string) (domain.URL, error) {
	var url domain.URL

	res := r.coll.FindOne(ctx, bson.M{
		"hash": hash,
	})
	if err := res.Err(); err != nil {
		return domain.URL{}, err
	}

	if err := res.Decode(&url); err != nil {
		return domain.URL{}, err
	}

	return url, nil
}
