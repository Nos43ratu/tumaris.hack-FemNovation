package repository

import (
	"database/sql"
	"io/ioutil"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/repository/auth"
	"tumaris.hack-FemNovation/backend/internal/repository/order"
	"tumaris.hack-FemNovation/backend/internal/repository/products"
	"tumaris.hack-FemNovation/backend/internal/repository/token"
	"tumaris.hack-FemNovation/backend/internal/repository/users"
)

type Repository struct {
	Auth     auth.Auth
	Token    token.Token
	Products products.Products
	Order    order.Order
	Users    users.Users
}

func New(db *pgxpool.Pool, sqlite *sql.DB, sqliteTimeout time.Duration, logger *zap.SugaredLogger) *Repository {
	dbTimeout := 10 * time.Second

	p1, err := ioutil.ReadFile("auth-private.pem")
	if err != nil {
		log.Println("error reading cert")
		return nil
	}

	p2, err := ioutil.ReadFile("auth-public.pem")
	if err != nil {
		log.Println("error reading cert")
		return nil
	}

	return &Repository{
		Auth:     auth.NewAuthRepo(logger, sqlite, sqliteTimeout),
		Order:    order.NewOrderRepo(logger, db, dbTimeout),
		Products: products.NewProductsRepo(logger, db, sqliteTimeout),
		Token:    token.NewTokenRepo(logger, sqlite, sqliteTimeout, p1, p2),
		Users:    users.NewUserRepo(logger, db, dbTimeout),
	}
}
