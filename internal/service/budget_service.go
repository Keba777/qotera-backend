package service

import (
	"context"
	"qotera-backend/internal/domain"
	"qotera-backend/internal/repository"
	"time"

	"github.com/google/uuid"
)

type BudgetService interface {
	SetBudget(ctx context.Context, budget *domain.Budget) error
	GetBudgets(ctx context.Context, userID uuid.UUID) ([]domain.Budget, error)
	CheckBudget(ctx context.Context, userID uuid.UUID, categoryID uuid.UUID, amount float64) (bool, error)
}

type budgetServiceImpl struct {
	budgetRepo      repository.BudgetRepository
	transactionRepo repository.TransactionRepository
}

func NewBudgetService(budgetRepo repository.BudgetRepository, transactionRepo repository.TransactionRepository) BudgetService {
	return &budgetServiceImpl{
		budgetRepo:      budgetRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *budgetServiceImpl) SetBudget(ctx context.Context, budget *domain.Budget) error {
	// Look for existing budget for this user and category to update, or create new
	existing, _ := s.budgetRepo.GetByUserID(ctx, budget.UserID)
	for _, b := range existing {
		if b.CategoryID == budget.CategoryID && b.Period == budget.Period {
			budget.ID = b.ID
			return s.budgetRepo.Update(ctx, budget)
		}
	}
	return s.budgetRepo.Create(ctx, budget)
}

func (s *budgetServiceImpl) GetBudgets(ctx context.Context, userID uuid.UUID) ([]domain.Budget, error) {
	return s.budgetRepo.GetByUserID(ctx, userID)
}

func (s *budgetServiceImpl) CheckBudget(ctx context.Context, userID uuid.UUID, categoryID uuid.UUID, amount float64) (bool, error) {
	// Simple implementation for now: Check if total spent in the current month exceeds limit
	now := time.Now()
	summary, err := s.transactionRepo.GetMonthlySummary(ctx, userID, now.Year(), int(now.Month()))
	if err != nil {
		return false, err
	}

	budgets, err := s.budgetRepo.GetByUserID(ctx, userID)
	if err != nil {
		return false, err
	}

	for _, b := range budgets {
		if b.CategoryID == categoryID && b.Period == domain.BudgetPeriodMonthly {
			if summary.TotalExpense+amount > b.LimitAmount {
				return true, nil
			}
		}
	}

	return false, nil
}
