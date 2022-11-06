package order

import (
	"context"
	"database/sql"
	"errors"
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

func (o *OrderRepo) GetAll(userID string) ([]*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), o.timeout)
	defer cancel()

	role := ""
	var sID sql.NullInt64
	query := `SELECT role, shop_id FROM users WHERE id=$1`
	err := o.db.QueryRow(ctx, query, userID).Scan(&role, &sID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			o.logger.Errorf("invite does not exist: %s", err)
			return nil, models.ErrInviteDoesNotExist
		}
		o.logger.Errorf("db error: %s", err)
		return nil, models.ErrDBConnection
	}

	if role == "client" {
		query = `SELECT * FROM orders WHERE client_id=$1`

		rows, err := o.db.Query(ctx, query, userID)
		if err != nil {
			o.logger.Errorf("db error: %s", err)
			return nil, models.ErrDBConnection
		}
		defer rows.Close()

		var orders []*models.Order
		for rows.Next() {
			var order *models.Order

			reason := &sql.NullString{}
			err = rows.Scan(&order.ID, &order.Status, &order.ClientID, &order.ShopID, &order.ProductID, &order.CreatedAt, &order.PayedAt, &order.PackedAt, &order.DeliveredAt, reason)
			if err != nil {
				o.logger.Errorf("db error: %s", err)
				return nil, models.ErrDBConnection
			}

			if reason.Valid {
				order.CancelReason = reason.String
			}
			orders = append(orders, order)
		}

		if err := rows.Err(); err != nil {
			o.logger.Errorf("db error: %s", err)
			return nil, models.ErrDBConnection
		}

		return orders, nil
	} else {
		query = `SELECT * FROM orders WHERE shop_id=$1`

		shopID := sID.Int64
		rows, err := o.db.Query(ctx, query, shopID)
		if err != nil {
			o.logger.Errorf("db error: %s", err)
			return nil, models.ErrDBConnection
		}
		defer rows.Close()

		orders := make([]*models.Order, 0, 20)

		for rows.Next() {
			order := &models.Order{}

			reason := &sql.NullString{}
			err = rows.Scan(&order.ID, &order.Status, &order.ClientID, &order.ShopID, &order.ProductID, &order.CreatedAt, &order.PayedAt, &order.PackedAt, &order.DeliveredAt, reason)
			if err != nil {
				o.logger.Errorf("db error: %s", err)
				return nil, models.ErrDBConnection
			}

			if reason.Valid {
				order.CancelReason = reason.String
			}

			orders = append(orders, order)
		}

		if err := rows.Err(); err != nil {
			o.logger.Errorf("db error: %s", err)
			return nil, models.ErrDBConnection
		}

		return orders, nil
	}
}

func (o *OrderRepo) GetByID(orderID string) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), o.timeout)
	defer cancel()

	order := &models.Order{}
	query := `SELECT * FROM orders WHERE id=$1`
	err := o.db.QueryRow(ctx, query, orderID).Scan(&order.ID, &order.Status, &order.ClientID, &order.ShopID, &order.ProductID, &order.CreatedAt, &order.PayedAt, &order.PackedAt, &order.DeliveredAt, &order.CancelReason)
	if err != nil {
		o.logger.Errorf("db error: %s", err)
		return nil, models.ErrDBConnection
	}

	return order, nil
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

func (o *OrderRepo) Update(orderID string, order *models.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), o.timeout)
	defer cancel()

	query := `UPDATE orders SET status=$1 WHERE id=$2`
	_, err := o.db.Exec(ctx, query, order.Status, orderID)
	if err != nil {
		o.logger.Errorf("db error: %s", err)
		return models.ErrDBConnection
	}

	query = `UPDATE orders SET cancel_reason=$1 WHERE id=$2`
	_, err = o.db.Exec(ctx, query, order.CancelReason, orderID)
	if err != nil {
		o.logger.Errorf("db error: %s", err)
		return models.ErrDBConnection
	}

	return nil
}
