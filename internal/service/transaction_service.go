package service

import (
	"context"
	"qotera-backend/internal/domain"

	"github.com/google/uuid"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction *domain.Transaction) error
	GetTransactions(ctx context.Context, userID uuid.UUID) ([]domain.Transaction, error)
	GetTransactionByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error)
}
