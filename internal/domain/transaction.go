package domain

import (
	"time"

	"github.com/google/uuid"
)

type TransactionType string
type TransactionSource string

const (
	TypeIncome  TransactionType = "income"
	TypeExpense TransactionType = "expense"

	SourceTelebirr TransactionSource = "telebirr"
	SourceCBE      TransactionSource = "cbe"
)

type Transaction struct {
	ID              uuid.UUID         `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID          uuid.UUID         `json:"user_id" gorm:"type:uuid;index;not null;uniqueIndex:idx_user_ref"`
	Amount          float64           `json:"amount" gorm:"type:decimal(15,2);not null"`
	Fee             *float64          `json:"fee" gorm:"type:decimal(15,2)"`
	Type            TransactionType   `json:"type" gorm:"type:varchar(20);not null"`
	Source          TransactionSource `json:"source" gorm:"type:varchar(50);not null"`
	ReferenceNumber *string           `json:"reference_number" gorm:"type:varchar(100);uniqueIndex:idx_user_ref"` // Composite unique index to avoid duplicate inserts
	TransactionDate time.Time         `json:"transaction_date" gorm:"not null;index"`
	RawMessage      string            `json:"raw_message" gorm:"type:text;not null"`
	CreatedAt       time.Time         `json:"created_at" gorm:"autoCreateTime"`
}
