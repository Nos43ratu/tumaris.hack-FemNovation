package service

import (
	"time"

	"tumaris.hack-FemNovation/backend/internal/repository"
	"tumaris.hack-FemNovation/backend/internal/service/auth"
	"tumaris.hack-FemNovation/backend/internal/service/order"
	"tumaris.hack-FemNovation/backend/pkg/hash"
)

type Service struct {
	Auth  auth.Service
	Order order.Service
}

func New(repos *repository.Repository, hasher *hash.BcryptHasher, accessTTL time.Duration, refreshTTL time.Duration) *Service {
	return &Service{
		Auth:  auth.NewAuthService(repos.Auth, hasher, repos.Token, accessTTL, refreshTTL),
		Order: order.NewOrderService(repos.Order),
	}
}
