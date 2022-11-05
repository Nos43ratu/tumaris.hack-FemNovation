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

func (o *OrderRepo) GetAll() ([]*models.Order, error) {
	return nil, nil
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

	if err = tx.Commit(ctx); err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			o.logger.Errorf("transaction error: %s", errTX)
		}
		o.logger.Errorf("db error: %s", err)
		return models.ErrDBConnection
	}

	return nil
}

func (o *OrderRepo) Update(order *models.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), o.timeout)
	defer cancel()

	tx, err := o.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		o.logger.Errorf("transaction error: %s", err)
		return models.ErrDBConnection
	}

	query := `UPDATE orders set status=$1 and cancel_reason=$2 WHERE id=$3`
	_, err = tx.Exec(ctx, query, order.Status, order.CancelReason, order.ID)
	if err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			o.logger.Errorf("transaction error: %s", errTX)
		}
		o.logger.Errorf("db error: %s", err)
		return models.ErrDBConnection
	}

	if err = tx.Commit(ctx); err != nil {
		errTX := tx.Rollback(ctx)
		if errTX != nil {
			o.logger.Errorf("transaction error: %s", errTX)
		}
		o.logger.Errorf("db error: %s", err)
		return models.ErrDBConnection
	}

	return nil
}
