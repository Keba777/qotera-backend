package handler

import (
	"qotera-backend/internal/domain"
	"qotera-backend/internal/middleware"
	"qotera-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type BudgetHandler struct {
	budgetService service.BudgetService
}

func NewBudgetHandler(budgetService service.BudgetService) *BudgetHandler {
	return &BudgetHandler{budgetService: budgetService}
}

func (h *BudgetHandler) SetBudget(c *fiber.Ctx) error {
	userID, ok := c.Locals(middleware.ContextKeyUserID).(uuid.UUID)
	if !ok {
		if fid := c.Locals("userID"); fid != nil {
			userID = fid.(uuid.UUID)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
	}

	var budget domain.Budget
	if err := c.BodyParser(&budget); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	budget.UserID = userID
	if err := h.budgetService.SetBudget(c.Context(), &budget); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set budget"})
	}

	return c.JSON(budget)
}

func (h *BudgetHandler) GetBudgets(c *fiber.Ctx) error {
	userID, ok := c.Locals(middleware.ContextKeyUserID).(uuid.UUID)
	if !ok {
		if fid := c.Locals("userID"); fid != nil {
			userID = fid.(uuid.UUID)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
	}

	budgets, err := h.budgetService.GetBudgets(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch budgets"})
	}

	return c.JSON(budgets)
}
