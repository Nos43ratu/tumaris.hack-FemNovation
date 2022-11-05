package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type AuthRepo struct {
	db      *sql.DB
	logger  *zap.SugaredLogger
	timeout time.Duration
}

func NewAuthRepo(logger *zap.SugaredLogger, db *sql.DB, timeout time.Duration) Auth {
	return &AuthRepo{
		db:      db,
		logger:  logger,
		timeout: timeout,
	}
}

func (a *AuthRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	user := &models.User{}
	query := `SELECT * FROM users where email = $1`
	err := a.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		a.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}

func (a *AuthRepo) GetUserByID(ID int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), a.timeout)
	defer cancel()

	user := &models.User{}
	query := `SELECT * FROM users where id = $1`
	err := a.db.QueryRowContext(ctx, query, ID).Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		a.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}
