package products

import "tumaris.hack-FemNovation/backend/internal/models"

type Products interface {
	CreateProduct(product *models.Product) (int, error)
	// GetUserByID(ID int) (*models.User, error)
}
