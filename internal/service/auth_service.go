package service

import (
	"context"
	"qotera-backend/internal/domain"
)

type AuthService interface {
	Register(ctx context.Context, user *domain.User, password string) error
	Login(ctx context.Context, identifier string, password string) (string, error) // Returns JWT token
}
