package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID                uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID            uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	ProviderID        uuid.UUID `json:"provider_id" gorm:"type:uuid;not null"`
	AccountIdentifier string    `json:"account_identifier" gorm:"not null"`
	Currency          string    `json:"currency" gorm:"default:'ETB'"`
	CreatedAt         time.Time `json:"created_at" gorm:"autoCreateTime"`
}
