package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Phone        string    `json:"phone" gorm:"uniqueIndex;not null"`
	Email        *string   `json:"email" gorm:"uniqueIndex"`
	PasswordHash string    `json:"-" gorm:"not null"`
	IsPremium    bool      `json:"is_premium" gorm:"default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
