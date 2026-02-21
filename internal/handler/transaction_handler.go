package handler

import (
	"qotera-backend/internal/domain"
	"qotera-backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

// SyncTransactions allows the mobile app to post an array of parsed SMS transactions safely
func (h *TransactionHandler) SyncTransactions(c *fiber.Ctx) error {
	// TODO: Replace with authenticated user ID from Middleware JWT Context
	// For testing purposes, we extract it from the header or mock it
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing User ID"})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
	}

	var transactions []domain.Transaction
	if err := c.BodyParser(&transactions); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON payload"})
	}

	if err := h.transactionService.SyncTransactions(c.Context(), userID, transactions); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to sync transactions"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Transactions synced successfully"})
}

// GetSummary returns daily/weekly/monthly analytics
func (h *TransactionHandler) GetSummary(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing User ID"})
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid User ID"})
	}

	timeframe := c.Query("timeframe", "monthly") // daily, weekly, monthly
	year, _ := strconv.Atoi(c.Query("year", "2026"))

	switch timeframe {
	case "daily":
		month, _ := strconv.Atoi(c.Query("month", "2"))
		day, _ := strconv.Atoi(c.Query("day", "21"))
		summary, err := h.transactionService.GetDailySummary(c.Context(), userID, year, month, day)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(summary)

	case "weekly":
		week, _ := strconv.Atoi(c.Query("week", "8"))
		summary, err := h.transactionService.GetWeeklySummary(c.Context(), userID, year, week)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(summary)

	case "monthly":
		month, _ := strconv.Atoi(c.Query("month", "2"))
		summary, err := h.transactionService.GetMonthlySummary(c.Context(), userID, year, month)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(summary)

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid timeframe specified"})
	}
}
