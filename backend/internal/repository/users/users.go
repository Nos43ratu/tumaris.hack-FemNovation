package users

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type UserRepo struct {
	db      *pgxpool.Pool
	logger  *zap.SugaredLogger
	timeout time.Duration
}

func NewUserRepo(logger *zap.SugaredLogger, db *pgxpool.Pool, timeout time.Duration) Users {
	return &UserRepo{
		db:      db,
		logger:  logger,
		timeout: timeout,
	}
}

func (u *UserRepo) GetByEmail(email string) (*models.UserInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	user := &models.UserInfo{}
	query := `SELECT * FROM users WHERE email=$1`
	log.Printf(">>[%s]:[%s]", query, email)
	err := u.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Phone, &user.FirstName, &user.LastName, &user.Password, &user.Role, &user.AboutMe, &user.Instagram, &user.Rating, &user.ShopID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		u.logger.Errorf("db error: %s", err)
		return nil, models.ErrDBConnection
	}

	return user, nil
}
