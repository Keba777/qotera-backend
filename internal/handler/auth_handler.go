package handler

import (
	"qotera-backend/internal/service"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	return c.SendString("Register endpoint")
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	return c.SendString("Login endpoint")
}
