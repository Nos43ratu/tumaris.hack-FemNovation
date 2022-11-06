package products

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"tumaris.hack-FemNovation/backend/internal/models"
)

type ProductsRepo struct {
	db      *pgxpool.Pool
	logger  *zap.SugaredLogger
	timeout time.Duration
}

func NewProductsRepo(logger *zap.SugaredLogger, db *pgxpool.Pool, timeout time.Duration) Products {
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

	err := r.db.QueryRow(ctx, "INSERT INTO product (shop_id, name, description, sizes, colors, weight, price, rating, category_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id",
		product.ShopID, product.Name, product.Description, pq.Array(product.Sizes), pq.Array(product.Colors), product.Weight, product.Price, 9.1, product.CategoryID).Scan(&id)
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

	err := r.db.QueryRow(ctx, query, product.Name, product.Description, pq.Array(product.Sizes), pq.Array(product.Colors), product.Weight, product.Price, product.CategoryID, product.ProductID).Scan(&id)
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
	query := `SELECT id, shop_id, name, description, sizes, colors, weight, price, rating, category_id FROM product where id = $1`
	err := r.db.QueryRow(ctx, query, ID).Scan(&res.ProductID, &res.ShopID, &res.Name, &res.Description, pq.Array(&res.Sizes), &res.Colors, &res.Weight, &res.Price, &res.Rating, &res.CategoryID)
	if err != nil {
		r.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, err
	}
	return res, nil
}

func (r *ProductsRepo) DeleteProduct(ID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	var res int
	query := `DELETE from product where id = $1 returning id`
	err := r.db.QueryRow(ctx, query, ID).Scan(&res)
	if err != nil {
		r.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return models.ErrNoRows
		}
		return err
	}

	return nil
}

func (r *ProductsRepo) GetProductsByCategory(ID int) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	var products []*models.Product

	query := `SELECT id, shop_id, name, description, sizes, colors, weight, price, rating, category_id FROM product where category_id = $1`
	rows, err := r.db.Query(ctx, query, ID)
	// Scan(&res.ProductID, &res.ShopID, &res.Name, &res.Description, pq.Array(&res.Sizes), &res.Colors, &res.Weight, &res.Price, &res.Rating, &res.CategoryID)
	if err != nil {
		r.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, err
	}
	for rows.Next() {
		res := &models.Product{}
		err := rows.Scan(&res.ProductID, &res.ShopID, &res.Name, &res.Description, pq.Array(&res.Sizes), &res.Colors, &res.Weight, &res.Price, &res.Rating, &res.CategoryID)
		if err != nil {
			r.logger.Errorf("db error: %s", err)
			return nil, models.ErrDBConnection
		}
		products = append(products, res)
	}
	return products, nil
}

func (r *ProductsRepo) GetProducts(filter *models.ProductFilter) ([]*models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	var products []*models.Product

	query := `SELECT id, shop_id, name, description, sizes, colors, weight, price, rating, category_id FROM product`

	if filter.Search != "" {
		query = query + fmt.Sprintf(" where (name LIKE '%%%s%%' OR description LIKE '%%%s%%')", filter.Search, filter.Search)
	}
	if filter.ShopID != "" {
		if filter.Search != "" {
			query = query + fmt.Sprintf(" and shop_id = %s", filter.ShopID)
		} else {
			query = query + fmt.Sprintf(" where shop_id = %s", filter.ShopID)
		}

	}

	fmt.Println(query)

	rows, err := r.db.Query(ctx, query)
	// Scan(&res.ProductID, &res.ShopID, &res.Name, &res.Description, pq.Array(&res.Sizes), &res.Colors, &res.Weight, &res.Price, &res.Rating, &res.CategoryID)
	if err != nil {
		r.logger.Errorf("db error: %s", err)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRows
		}
		return nil, err
	}
	for rows.Next() {
		res := &models.Product{}
		err := rows.Scan(&res.ProductID, &res.ShopID, &res.Name, &res.Description, pq.Array(&res.Sizes), &res.Colors, &res.Weight, &res.Price, &res.Rating, &res.CategoryID)
		if err != nil {
			r.logger.Errorf("db error: %s", err)
			return nil, models.ErrDBConnection
		}
		products = append(products, res)
	}
	return products, nil
}
