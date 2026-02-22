package handler

import (
	"qotera-backend/internal/middleware"
	"qotera-backend/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authService service.AuthService
	userService service.UserService
}

func NewAuthHandler(authService service.AuthService, userService service.UserService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		userService: userService,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	return c.SendString("Register endpoint")
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Phone string `json:"phone"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Phone number is required"})
	}

	if err := h.authService.SendOTP(c.Context(), req.Phone); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to send OTP"})
	}

	return c.JSON(fiber.Map{"message": "OTP sent successfully"})
}

func (h *AuthHandler) VerifyOTP(c *fiber.Ctx) error {
	type VerifyRequest struct {
		Phone string `json:"phone"`
		OTP   string `json:"otp"`
	}

	var req VerifyRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.Phone == "" || req.OTP == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Phone and OTP are required"})
	}

	token, err := h.authService.VerifyOTP(c.Context(), req.Phone, req.OTP)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid OTP"})
	}

	return c.JSON(fiber.Map{"token": token})
}
func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID, ok := c.Locals(middleware.ContextKeyUserID).(uuid.UUID)
	if !ok {
		// Try legacy
		if fid := c.Locals("userID"); fid != nil {
			userID = fid.(uuid.UUID)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
	}

	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user"})
	}

	if user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}
