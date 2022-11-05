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

func (s *ProductsService) CreateProduct(product *models.Product) (int, error){
	return s.products.CreateProduct(product)
}