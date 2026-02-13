package domain

import (
	"time"

	"github.com/google/uuid"
)

type BudgetPeriod string

const (
	BudgetPeriodDaily   BudgetPeriod = "daily"
	BudgetPeriodWeekly  BudgetPeriod = "weekly"
	BudgetPeriodMonthly BudgetPeriod = "monthly"
)

type Budget struct {
	ID          uuid.UUID    `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID      uuid.UUID    `json:"user_id" gorm:"type:uuid;not null;index"`
	CategoryID  uuid.UUID    `json:"category_id" gorm:"type:uuid;not null;index"`
	LimitAmount float64      `json:"limit_amount" gorm:"not null"`
	Period      BudgetPeriod `json:"period" gorm:"type:varchar(20);not null"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
}
