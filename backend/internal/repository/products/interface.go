package products

import "tumaris.hack-FemNovation/backend/internal/models"

type Products interface {
	CreateProduct(product *models.Product) (int, error)
	UpdateProduct(product *models.Product) (int, error)
	GetProductByID(ID int) (*models.Product, error)
	DeleteProduct(ID int) (error)
	GetProductsByCategory(ID int) ([]*models.Product, error)
	GetProducts() ([]*models.Product, error)
}
