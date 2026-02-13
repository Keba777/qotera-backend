package repository

import (
	"context"

	"qotera-backend/internal/domain"

	"github.com/google/uuid"
)

type BudgetRepository interface {
	Create(ctx context.Context, budget *domain.Budget) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Budget, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Budget, error)
	Update(ctx context.Context, budget *domain.Budget) error
	Delete(ctx context.Context, id uuid.UUID) error
}
