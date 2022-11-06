package products

import (
	// "errors"
	// "fmt"
	// "time"

	// "github.com/dgrijalva/jwt-go"
	// "github.com/google/uuid"
	// "golang.org/x/crypto/bcrypt"

	"tumaris.hack-FemNovation/backend/internal/models"
	"tumaris.hack-FemNovation/backend/internal/repository/products"
	// "tumaris.hack-FemNovation/backend/internal/repository/token"
	// "tumaris.hack-FemNovation/backend/pkg/hash"
)

type ProductsService struct {
	products products.Products
}

func NewProductsService(products products.Products) Service {
	return &ProductsService{
		products: products,
	}
}

func (s *ProductsService) CreateProduct(product *models.Product) (int, error) {
	return s.products.CreateProduct(product)
}

func (s *ProductsService) UpdateProduct(product *models.Product) (int, error) {
	initial, err := s.products.GetProductByID(product.ProductID)
	if err != nil {
		return -1, err
	}
	if product.Name == "" {
		product.Name = initial.Name
	}
	if product.Description == "" {
		product.Description = initial.Description
	}
	if product.Sizes == nil {
		product.Sizes = initial.Sizes
	}
	if product.Colors == nil {
		product.Colors = initial.Colors
	}
	if product.Weight == 0 {
		product.Weight = initial.Weight
	}
	if product.Price == 0 {
		product.Price = initial.Price
	}
	if product.CategoryID == 0 {
		product.CategoryID = initial.CategoryID
	}
	return s.products.UpdateProduct(product)
}

func (s *ProductsService) DeleteProduct(ID int) error {
	return s.products.DeleteProduct(ID)
}

func (s *ProductsService) GetProduct(ID int) (*models.Product, error) {
	return s.products.GetProductByID(ID)
}

func (s *ProductsService) GetProductsByCategory(ID int) ([]*models.Product, error) {
	return s.products.GetProductsByCategory(ID)
}

func (s *ProductsService) GetProducts() ([]*models.Product, error) {
	return s.products.GetProducts()
}