package service

import (
	"telegram-bot/internal/models"
	"telegram-bot/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}

// Пример метода сервиса
func (s *Service) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
