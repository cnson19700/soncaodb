package repository

import (
	"context"

	"github.com/soncaodb/repository/student"
	"gorm.io/gorm"
)

type Repository struct {
	Student student.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		Student: student.NewPGRepository(getSQLClient),
	}
}
