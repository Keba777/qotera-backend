package handler

import (
	"qotera-backend/internal/service"

	"github.com/gofiber/fiber/v3"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) CreateTransaction(c fiber.Ctx) error {
	return c.SendString("Create Transaction endpoint")
}

func (h *TransactionHandler) GetTransactions(c fiber.Ctx) error {
	return c.SendString("Get Transactions endpoint")
}
