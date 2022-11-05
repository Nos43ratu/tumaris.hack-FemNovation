package order

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type OrderRepo struct {
	db      *pgxpool.Pool
	logger  *zap.SugaredLogger
	timeout time.Duration
}

func NewOrderRepo(logger *zap.SugaredLogger, db *pgxpool.Pool, timeout time.Duration) Order {
	return &OrderRepo{
		db:      db,
		logger:  logger,
		timeout: timeout,
	}
}

func (o *OrderRepo) Create(order *models.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), o.timeout)
	defer cancel()

	tx, err := o.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		o.logger.Errorf("transaction error: %s", err)
		return models.ErrDBConnection
	}

	query := `INSERT INTO orders (status, client_id, shop_id, product_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err = tx.QueryRow(ctx, query, 0, order.ClientID, order.ShopID, order.ProductID).Scan(&order.ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			o.logger.Errorf("transaction error: %s", errTX)
		}
		o.logger.Errorf("db error: %s", err)
		return models.ErrDBConnection
	}

	return nil
}
