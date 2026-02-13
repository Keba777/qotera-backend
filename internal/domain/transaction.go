package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string

const (
	TransactionTypeCredit TransactionType = "credit"
	TransactionTypeDebit  TransactionType = "debit"
)

type Transaction struct {
	ID              uuid.UUID       `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID          uuid.UUID       `json:"user_id" gorm:"type:uuid;not null;index"`
	AccountID       uuid.UUID       `json:"account_id" gorm:"type:uuid;not null;index"`
	Amount          float64         `json:"amount" gorm:"not null"`
	Currency        string          `json:"currency" gorm:"default:'ETB'"`
	TransactionType TransactionType `json:"transaction_type" gorm:"type:varchar(20);not null"`
	CategoryID      *uuid.UUID      `json:"category_id" gorm:"type:uuid;index"`
	RawMessage      string          `json:"raw_message" gorm:"type:text"`
	Reference       *string         `json:"reference"`
	BalanceAfter    *float64        `json:"balance_after"`
	TransactionDate time.Time       `json:"transaction_date" gorm:"not null"`
	CreatedAt       time.Time       `json:"created_at" gorm:"autoCreateTime"`
}
