package service

import (
	"context"
	"errors"
	"fmt"
	"qotera-backend/internal/domain"
	"qotera-backend/internal/repository"
	"qotera-backend/pkg/utils"
)

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (s *authServiceImpl) Login(ctx context.Context, identifier string, password string) (string, error) {
	// Old implementation logic, might be unused or for email/password
	return "", errors.New("not implemented")
}

func (s *authServiceImpl) Register(ctx context.Context, user *domain.User, password string) error {
	// Old implementation logic
	return errors.New("not implemented")
}

func (s *authServiceImpl) SendOTP(ctx context.Context, phone string) error {
	// In a real app, send SMS here.
	fmt.Printf("Sending OTP for %s: %s\n", phone, "123456")
	return nil
}

func (s *authServiceImpl) VerifyOTP(ctx context.Context, phone string, otp string) (string, error) {
	if otp != "123456" {
		return "", errors.New("invalid OTP")
	}

	user, err := s.userRepo.GetByPhone(ctx, phone)
	if err != nil {
		// Real DB error
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	if user == nil {
		// User not found — auto-register on first OTP verify
		user = &domain.User{
			Phone: phone,
		}
		if createErr := s.userRepo.Create(ctx, user); createErr != nil {
			return "", fmt.Errorf("failed to create user: %w", createErr)
		}
	}

	// Generate JWT
	token, err := utils.GenerateToken(user.ID.String(), user.Phone)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}
