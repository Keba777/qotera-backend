package repository

import (
	"context"
	"errors"
	"qotera-backend/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type budgetRepositoryImpl struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) BudgetRepository {
	return &budgetRepositoryImpl{db: db}
}

func (r *budgetRepositoryImpl) Create(ctx context.Context, budget *domain.Budget) error {
	return r.db.WithContext(ctx).Create(budget).Error
}

func (r *budgetRepositoryImpl) GetByID(ctx context.Context, id uuid.UUID) (*domain.Budget, error) {
	var budget domain.Budget
	if err := r.db.WithContext(ctx).First(&budget, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &budget, nil
}

func (r *budgetRepositoryImpl) GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Budget, error) {
	var budgets []domain.Budget
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		return nil, err
	}
	return budgets, nil
}

func (r *budgetRepositoryImpl) Update(ctx context.Context, budget *domain.Budget) error {
	return r.db.WithContext(ctx).Save(budget).Error
}

func (r *budgetRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Budget{}, "id = ?", id).Error
}
