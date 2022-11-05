package products

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type ProductsRepo struct {
	db      *sql.DB
	logger  *zap.SugaredLogger
	timeout time.Duration
}

func NewProductsRepo(logger *zap.SugaredLogger, db *sql.DB, timeout time.Duration) Products {
	return &ProductsRepo{
		db:      db,
		logger:  logger,
		timeout: timeout,
	}
}

func (r *ProductsRepo) CreateProduct(product *models.Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	var id int
	defer cancel()

	err := r.db.QueryRowContext(ctx, "INSERT INTO product (shop_id, name, description, sizes, colors, weight, price, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7)",
	product.ShopID, product.Name, product.Description, pq.Array(product.Sizes), pq.Array(product.Colors), product.Weight, product.Price, product.CategoryID).Scan(&id)
	if err != nil {
		r.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return -1, models.ErrNoRows
		}
		return -1, err
	}

	return id, nil
}

func (r *ProductsRepo) UpdateProduct(product *models.Product) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	var id int
	defer cancel()

	query := `UPDATE product SET 
	name = $1,
	description = $2, 
	sizes = $3,
	colors = $4, 
	weight = $5,
	price = $6,
	category_id = $7
		where id = $8
	returning id`

	err := r.db.QueryRowContext(ctx, query, product.Name, product.Description, pq.Array(product.Sizes), pq.Array(product.Colors), product.Weight, product.Price, product.CategoryID, product.ProductID).Scan(&id)
	if err != nil {
		r.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			r.logger.Errorf("db error: %s", err)
			return -1, models.ErrNoRows
		}
		return -1, err
	}

	return id, nil
}

func (r *ProductsRepo) GetProductByID(ID int) (*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	res := &models.Product{}
	query := `SELECT * FROM products where id = $1`
	err := r.db.QueryRowContext(ctx, query, ID).Scan(&res.ProductID, &res.ShopID, &res.Name, &res.Description, pq.Array(&res.Sizes), pq.Array(&res.Colors), &res.Weight, &res.Price, &res.Rating, &res.CategoryID)
	if err != nil {
		r.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, err
	}

	return res, nil
}