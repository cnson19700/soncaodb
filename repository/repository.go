package repository

import (
	"context"

	"github.com/soncaodb/repository/book"
	"github.com/soncaodb/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	User user.Repository
	Book book.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		User: user.NewPGRepository(getSQLClient),
		Book: book.NewPGRepository(getSQLClient),
	}
}
