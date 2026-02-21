package service_test

import (
	"context"
	"errors"
	"testing"

	"qotera-backend/internal/domain"
	"qotera-backend/internal/service"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ─── Mock User Repository ────────────────────────────────────────────────────

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	// Simulate database assigning a UUID on create
	if args.Error(0) == nil && user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetByPhone(ctx context.Context, phone string) (*domain.User, error) {
	args := m.Called(ctx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// ─── Tests ───────────────────────────────────────────────────────────────────

func TestSendOTP_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authSvc := service.NewAuthService(mockRepo)

	err := authSvc.SendOTP(context.Background(), "+251911234567")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestVerifyOTP_WrongOTP_ReturnsError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authSvc := service.NewAuthService(mockRepo)

	token, err := authSvc.VerifyOTP(context.Background(), "+251911234567", "000000")

	assert.Error(t, err)
	assert.Equal(t, "invalid OTP", err.Error())
	assert.Empty(t, token)
	mockRepo.AssertNotCalled(t, "GetByPhone")
}

func TestVerifyOTP_NewUser_CreatesUserAndReturnsToken(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authSvc := service.NewAuthService(mockRepo)

	phone := "+251911111111"

	// GetByPhone returns nil, nil → user does not yet exist
	mockRepo.On("GetByPhone", mock.Anything, phone).Return(nil, nil)
	// Create is called, mock simulates assigning a UUID
	mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
		return u.Phone == phone
	})).Return(nil)

	token, err := authSvc.VerifyOTP(context.Background(), phone, "123456")

	assert.NoError(t, err)
	assert.NotEmpty(t, token, "should return a JWT token")
	mockRepo.AssertExpectations(t)
}

func TestVerifyOTP_ExistingUser_ReturnsToken(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authSvc := service.NewAuthService(mockRepo)

	phone := "+251911222222"
	existingUser := &domain.User{
		ID:    uuid.New(),
		Phone: phone,
	}

	// User already exists
	mockRepo.On("GetByPhone", mock.Anything, phone).Return(existingUser, nil)
	// Create should NOT be called

	token, err := authSvc.VerifyOTP(context.Background(), phone, "123456")

	assert.NoError(t, err)
	assert.NotEmpty(t, token, "should return a JWT token for existing user")
	mockRepo.AssertNotCalled(t, "Create")
	mockRepo.AssertExpectations(t)
}

func TestVerifyOTP_DBError_ReturnsError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authSvc := service.NewAuthService(mockRepo)

	phone := "+251911333333"
	dbErr := errors.New("connection refused")

	mockRepo.On("GetByPhone", mock.Anything, phone).Return(nil, dbErr)

	token, err := authSvc.VerifyOTP(context.Background(), phone, "123456")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get user")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestVerifyOTP_CreateFails_ReturnsError(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authSvc := service.NewAuthService(mockRepo)

	phone := "+251911444444"
	createErr := errors.New("unique constraint violation")

	mockRepo.On("GetByPhone", mock.Anything, phone).Return(nil, nil)
	mockRepo.On("Create", mock.Anything, mock.Anything).Return(createErr)

	token, err := authSvc.VerifyOTP(context.Background(), phone, "123456")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create user")
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}
