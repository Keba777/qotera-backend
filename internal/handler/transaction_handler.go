package handler

import (
	"log"
	"qotera-backend/internal/domain"
	"qotera-backend/internal/middleware"
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
	userID, ok := c.Locals(middleware.ContextKeyUserID).(uuid.UUID)
	if !ok {
		// Try legacy
		if fid := c.Locals("userID"); fid != nil {
			userID = fid.(uuid.UUID)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
	}

	var transactions []domain.Transaction
	if err := c.BodyParser(&transactions); err != nil {
		log.Printf("SyncTransactions: BodyParser failed: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON payload: " + err.Error()})
	}

	if err := h.transactionService.SyncTransactions(c.Context(), userID, transactions); err != nil {
		log.Printf("SyncTransactions: Service call failed: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to sync transactions: " + err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Transactions synced successfully"})
}

// GetSummary returns daily/weekly/monthly analytics
func (h *TransactionHandler) GetSummary(c *fiber.Ctx) error {
	userID, ok := c.Locals(middleware.ContextKeyUserID).(uuid.UUID)
	if !ok {
		// Try legacy
		if fid := c.Locals("userID"); fid != nil {
			userID = fid.(uuid.UUID)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
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

// GetTransactions returns a list of historical transactions
func (h *TransactionHandler) GetTransactions(c *fiber.Ctx) error {
	userID, ok := c.Locals(middleware.ContextKeyUserID).(uuid.UUID)
	if !ok {
		// Try legacy
		if fid := c.Locals("userID"); fid != nil {
			userID = fid.(uuid.UUID)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
	}

	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	transactions, err := h.transactionService.GetTransactions(c.Context(), userID, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch transactions"})
	}

	return c.JSON(transactions)
}
