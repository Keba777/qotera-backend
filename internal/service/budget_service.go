package service

import (
	"context"
	"qotera-backend/internal/domain"

	"github.com/google/uuid"
)

type BudgetService interface {
	SetBudget(ctx context.Context, budget *domain.Budget) error
	GetBudgets(ctx context.Context, userID uuid.UUID) ([]domain.Budget, error)
	CheckBudget(ctx context.Context, userID uuid.UUID, categoryID uuid.UUID, amount float64) (bool, error) // Returns true if budget exceeded
}
