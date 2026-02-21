package handler

import (
	"qotera-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AnalyticsHandler struct {
	analyticsService service.AnalyticsService
}

func NewAnalyticsHandler(analyticsService service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{analyticsService: analyticsService}
}

func (h *AnalyticsHandler) GetMonthlySpending(c *fiber.Ctx) error {
	return c.SendString("Get Monthly Spending endpoint")
}
