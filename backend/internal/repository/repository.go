package repository

import (
	"github.com/go-redis/redis/v7"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Repository struct {
}

func New(db *pgxpool.Pool, rdb *redis.Client, sugar *zap.SugaredLogger) *Repository {
	return &Repository{}
}
