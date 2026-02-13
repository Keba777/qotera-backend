package service

import (
	"context"

	"github.com/google/uuid"
)

type AnalyticsService interface {
	GetMonthlySpending(ctx context.Context, userID uuid.UUID, month int, year int) (float64, error)
	GetCategoryBreakdown(ctx context.Context, userID uuid.UUID, month int, year int) (map[string]float64, error)
}
