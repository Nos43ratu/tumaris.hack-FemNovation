package products

import "tumaris.hack-FemNovation/backend/internal/models"

type Products interface {
	CreateProduct(product *models.Product) (int, error)
	UpdateProduct(product *models.Product) (int, error)
	GetProductByID(ID int) (*models.Product, error)
	DeleteProduct(ID int) (error)
	// GetUserByID(ID int) (*models.User, error)
}
