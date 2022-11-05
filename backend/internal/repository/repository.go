package repository

import (
	"database/sql"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/repository/auth"
	"tumaris.hack-FemNovation/backend/internal/repository/token"
)

type Repository struct {
	Auth  auth.Auth
	Token token.Token
}

func New(db *pgxpool.Pool, sqlite *sql.DB, sqliteTimeout time.Duration, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		Auth: auth.NewAuthRepo(logger, sqlite, sqliteTimeout),
	}
}
