package repository

import (
	"context"

	"qotera-backend/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionSummary struct {
	TotalIncome  float64 `json:"total_income"`
	TotalExpense float64 `json:"total_expense"`
	Balance      float64 `json:"balance"`
}

type TransactionRepository interface {
	BulkInsertTransactions(ctx context.Context, transactions []domain.Transaction) error
	GetMonthlySummary(ctx context.Context, userID uuid.UUID, year int, month int) (*TransactionSummary, error)
	GetWeeklySummary(ctx context.Context, userID uuid.UUID, year int, week int) (*TransactionSummary, error)
	GetDailySummary(ctx context.Context, userID uuid.UUID, year int, month int, day int) (*TransactionSummary, error)
	GetTransactions(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]domain.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

// BulkInsertTransactions inserts an array of transactions. If duplicate reference numbers match the user ID, it simply ignores it instead of erroring out.
func (r *transactionRepository) BulkInsertTransactions(ctx context.Context, transactions []domain.Transaction) error {
	// GORM Clause handles Postgres ON CONFLICT efficiently
	return r.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "reference_number"}},
		DoNothing: true,
	}).Create(&transactions).Error
}

// GetMonthlySummary returns aggregated totals for the specified month
func (r *transactionRepository) GetMonthlySummary(ctx context.Context, userID uuid.UUID, year int, month int) (*TransactionSummary, error) {
	var summary TransactionSummary

	err := r.db.WithContext(ctx).
		Table("transactions").
		Select("COALESCE(SUM(CASE WHEN type = ? THEN amount ELSE 0 END), 0) as total_income, COALESCE(SUM(CASE WHEN type = ? THEN amount + COALESCE(fee, 0) ELSE 0 END), 0) as total_expense", domain.TypeIncome, domain.TypeExpense).
		Where("user_id = ? AND EXTRACT(YEAR FROM transaction_date) = ? AND EXTRACT(MONTH FROM transaction_date) = ?", userID, year, month).
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	summary.Balance = summary.TotalIncome - summary.TotalExpense
	return &summary, nil
}

// GetWeeklySummary returns aggregated totals for a given year and week number
func (r *transactionRepository) GetWeeklySummary(ctx context.Context, userID uuid.UUID, year int, week int) (*TransactionSummary, error) {
	var summary TransactionSummary

	err := r.db.WithContext(ctx).
		Table("transactions").
		Select("COALESCE(SUM(CASE WHEN type = ? THEN amount ELSE 0 END), 0) as total_income, COALESCE(SUM(CASE WHEN type = ? THEN amount + COALESCE(fee, 0) ELSE 0 END), 0) as total_expense", domain.TypeIncome, domain.TypeExpense).
		Where("user_id = ? AND EXTRACT(IYYY FROM transaction_date) = ? AND EXTRACT(IW FROM transaction_date) = ?", userID, year, week).
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	summary.Balance = summary.TotalIncome - summary.TotalExpense
	return &summary, nil
}

// GetDailySummary returns aggregated totals for a specific day
func (r *transactionRepository) GetDailySummary(ctx context.Context, userID uuid.UUID, year int, month int, day int) (*TransactionSummary, error) {
	var summary TransactionSummary

	err := r.db.WithContext(ctx).
		Table("transactions").
		Select("COALESCE(SUM(CASE WHEN type = ? THEN amount ELSE 0 END), 0) as total_income, COALESCE(SUM(CASE WHEN type = ? THEN amount + COALESCE(fee, 0) ELSE 0 END), 0) as total_expense", domain.TypeIncome, domain.TypeExpense).
		Where("user_id = ? AND EXTRACT(YEAR FROM transaction_date) = ? AND EXTRACT(MONTH FROM transaction_date) = ? AND EXTRACT(DAY FROM transaction_date) = ?", userID, year, month, day).
		Scan(&summary).Error

	if err != nil {
		return nil, err
	}

	summary.Balance = summary.TotalIncome - summary.TotalExpense
	return &summary, nil
}

func (r *transactionRepository) GetTransactions(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("transaction_date DESC").
		Limit(limit).
		Offset(offset).
		Find(&transactions).Error

	return transactions, err
}
