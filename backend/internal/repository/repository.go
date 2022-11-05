package repository

import (
	"database/sql"
	"io/ioutil"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/repository/auth"
	"tumaris.hack-FemNovation/backend/internal/repository/products"
	"tumaris.hack-FemNovation/backend/internal/repository/order"
	"tumaris.hack-FemNovation/backend/internal/repository/token"
)

type Repository struct {
	Auth  auth.Auth
	Products products.Products
	Token token.Token
	Order order.Order
}

func New(db *pgxpool.Pool, sqlite *sql.DB, sqliteTimeout time.Duration, logger *zap.SugaredLogger) *Repository {
	dbTimeout := 10 * time.Second

	p1, err := ioutil.ReadFile("./build/id_rsa")
	if err != nil {
		log.Println("error reading cert")
		return nil
	}

	p2, err := ioutil.ReadFile("./build/id_rsa.pub")
	if err != nil {
		log.Println("error reading cert")
		return nil
	}

	return &Repository{
		Auth: auth.NewAuthRepo(logger, sqlite, sqliteTimeout),
		Products: products.NewProductsRepo(logger, sqlite, sqliteTimeout),
		Order: order.NewOrderRepo(logger, db, dbTimeout),
		Token: token.NewTokenRepo(logger, sqlite, sqliteTimeout, p1, p2),
	}
}
