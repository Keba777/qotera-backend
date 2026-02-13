package handler

import (
	"qotera-backend/internal/service"

	"github.com/gofiber/fiber/v3"
)

type BudgetHandler struct {
	budgetService service.BudgetService
}

func NewBudgetHandler(budgetService service.BudgetService) *BudgetHandler {
	return &BudgetHandler{budgetService: budgetService}
}

func (h *BudgetHandler) SetBudget(c fiber.Ctx) error {
	return c.SendString("Set Budget endpoint")
}

func (h *BudgetHandler) GetBudgets(c fiber.Ctx) error {
	return c.SendString("Get Budgets endpoint")
}
