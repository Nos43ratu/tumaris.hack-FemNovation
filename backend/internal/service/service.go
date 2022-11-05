package service

import "tumaris.hack-FemNovation/backend/internal/repository"

type Service struct {
}

func New(repos *repository.Repository) *Service {
	return &Service{}
}
