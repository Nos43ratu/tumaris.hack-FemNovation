package service

import (
	"tumaris.hack-FemNovation/backend/internal/repository"
	"tumaris.hack-FemNovation/backend/internal/service/auth"
)

type Service struct {
	Auth auth.Service
}

func New(repos *repository.Repository) *Service {
	return &Service{}
}
