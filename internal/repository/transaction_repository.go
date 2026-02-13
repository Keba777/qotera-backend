package repository

import (
	"context"

	"qotera-backend/internal/domain"

	"github.com/google/uuid"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *domain.Transaction) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Transaction, error)
	GetByAccountID(ctx context.Context, accountID uuid.UUID) ([]domain.Transaction, error)
	Update(ctx context.Context, transaction *domain.Transaction) error
	Delete(ctx context.Context, id uuid.UUID) error
}
