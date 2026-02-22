package service

import (
	"context"
	"qotera-backend/internal/domain"
	"qotera-backend/internal/repository"

	"github.com/google/uuid"
)

type TransactionService interface {
	SyncTransactions(ctx context.Context, userID uuid.UUID, transactions []domain.Transaction) error
	GetMonthlySummary(ctx context.Context, userID uuid.UUID, year int, month int) (*repository.TransactionSummary, error)
	GetWeeklySummary(ctx context.Context, userID uuid.UUID, year int, week int) (*repository.TransactionSummary, error)
	GetDailySummary(ctx context.Context, userID uuid.UUID, year int, month int, day int) (*repository.TransactionSummary, error)
	GetTransactions(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]domain.Transaction, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) TransactionService {
	return &transactionService{repo: repo}
}

func (s *transactionService) SyncTransactions(ctx context.Context, userID uuid.UUID, transactions []domain.Transaction) error {
	// Ensure all incoming transactions are correctly attributed to the authenticated user
	for i := range transactions {
		transactions[i].UserID = userID
	}
	return s.repo.BulkInsertTransactions(ctx, transactions)
}

func (s *transactionService) GetMonthlySummary(ctx context.Context, userID uuid.UUID, year int, month int) (*repository.TransactionSummary, error) {
	return s.repo.GetMonthlySummary(ctx, userID, year, month)
}

func (s *transactionService) GetWeeklySummary(ctx context.Context, userID uuid.UUID, year int, week int) (*repository.TransactionSummary, error) {
	return s.repo.GetWeeklySummary(ctx, userID, year, week)
}

func (s *transactionService) GetDailySummary(ctx context.Context, userID uuid.UUID, year int, month int, day int) (*repository.TransactionSummary, error) {
	return s.repo.GetDailySummary(ctx, userID, year, month, day)
}

func (s *transactionService) GetTransactions(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]domain.Transaction, error) {
	return s.repo.GetTransactions(ctx, userID, limit, offset)
}
