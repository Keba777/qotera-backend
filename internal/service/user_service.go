package service

import (
	"context"
	"qotera-backend/internal/domain"
	"qotera-backend/internal/repository"

	"github.com/google/uuid"
)

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	GetByPhone(ctx context.Context, phone string) (*domain.User, error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

func (s *userServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userServiceImpl) GetByPhone(ctx context.Context, phone string) (*domain.User, error) {
	return s.userRepo.GetByPhone(ctx, phone)
}
