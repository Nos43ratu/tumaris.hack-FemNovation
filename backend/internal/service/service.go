package service

import (
	"time"

	"tumaris.hack-FemNovation/backend/internal/repository"
	"tumaris.hack-FemNovation/backend/internal/service/auth"
	"tumaris.hack-FemNovation/backend/internal/service/order"
	"tumaris.hack-FemNovation/backend/internal/service/products"
	"tumaris.hack-FemNovation/backend/internal/service/users"
	"tumaris.hack-FemNovation/backend/pkg/hash"
)

type Service struct {
	Auth     auth.Service
	Products products.Service
	Order    order.Service
	Users    users.Service
}

func New(repos *repository.Repository, hasher *hash.BcryptHasher, accessTTL time.Duration, refreshTTL time.Duration) *Service {
	return &Service{
		Auth:     auth.NewAuthService(repos.Auth, hasher, repos.Token, accessTTL, refreshTTL),
		Products: products.NewProductsService(repos.Products),
		Order:    order.NewOrderService(repos.Order, repos.Products),
		Users:    users.NewUserService(repos.Users),
	}
}
