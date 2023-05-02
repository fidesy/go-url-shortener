package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fidesy/go-url-shortener/internal/domain"
	"github.com/fidesy/go-url-shortener/internal/repository"
)

type User interface {
	Create(user domain.User) (interface{}, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (interface{}, error)
}

const (
	salt       = "sH6tPbN89ghtQk98"
	signingKey = "jhk!kj4378bvk3f98JBNMu3njnHK"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID interface{} `json:"user_id"`
}

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

var _ User = &UserService{}

func (s *UserService) Create(user domain.User) (interface{}, error) {
	exists, err := s.repo.User.UsernameExists(context.Background(), user.Username)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("username already in use")
	}

	user.Password = generatePasswordHash(user.Password)
	return s.repo.User.Create(context.Background(), user)
}

func (s *UserService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.User.Get(context.Background(), username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *UserService) ParseToken(accessToken string) (interface{}, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
