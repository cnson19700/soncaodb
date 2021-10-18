package repository

import (
	"context"

	"github.com/github.com/soncaodb/repository/comment"
	"github.com/github.com/soncaodb/repository/movie"
	"github.com/github.com/soncaodb/repository/user"
	"github.com/github.com/soncaodb/repository/userfavorite"
	"gorm.io/gorm"
)

type Repository struct {
	User      user.Repository
	Movie     movie.Repository
	UserFavor userfavorite.Repository
	Comment   comment.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		User:      user.NewPGRepository(getSQLClient),
		Movie:     movie.NewPGRepository(getSQLClient),
		UserFavor: userfavorite.NewPGRepository(getSQLClient),
		Comment:   comment.NewPGRepository(getSQLClient),
	}
}
