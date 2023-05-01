package mongo

import (
	"context"
	"errors"

	"github.com/fidesy/go-url-shortener/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(coll *mongo.Collection) *UserRepository {
	return &UserRepository{coll: coll}
}

var _ domain.UserRepository = &UserRepository{}

func (r *UserRepository) Create(ctx context.Context, user domain.User) (interface{}, error) {
	res, err := r.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (r *UserRepository) Get(ctx context.Context, username, password string) (domain.User, error) {
	var user domain.User

	res := r.coll.FindOne(ctx, bson.M{
		"username":      username,
		"password_hash": password,
	})
	if err := res.Err(); err != nil {
		return domain.User{}, err
	}

	if err := res.Decode(&user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user domain.User) error {
	filter := bson.M{
		"username":      user.Username,
		"password_hash": user.Password,
	}
	update := bson.M{"$set": bson.M{
		"name":          user.Name,
		"username":      user.Username,
		"password_hash": user.Password,
	}}

	_, err := r.coll.UpdateOne(ctx, filter, update)

	return err
}

func (r *UserRepository) Delete(ctx context.Context, username, password string) error {
	filter := bson.M{
		"username":      username,
		"password_hash": password,
	}

	_, err := r.coll.DeleteOne(ctx, filter)
	return err
}

func (r *UserRepository) UsernameExists(ctx context.Context, username string) (bool, error) {
	var user domain.User

	res := r.coll.FindOne(ctx, bson.M{
		"username": username,
	})
	if err := res.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}

		return true, err
	}
	if err := res.Decode(&user); err != nil {
		return false, err
	}

	return user.Username == username, nil
}
