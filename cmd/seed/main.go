package main

import (
	"log"
	"os"

	"qotera-backend/internal/domain"
	"qotera-backend/pkg/config"
	"qotera-backend/pkg/database"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env (or env vars from Render)
	_ = godotenv.Load()
	cfg := config.LoadConfig()

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}

	// Create a test user if not exists
	phone := os.Getenv("SEED_USER_PHONE")
	if phone == "" {
		phone = "+251912345678" // default test phone
	}

	// Check if user already exists
	var existing domain.User
	if err := db.First(&existing, "phone = ?", phone).Error; err == nil {
		log.Printf("User with phone %s already exists, skipping seed", phone)
		return
	}

	user := domain.User{
		ID:           uuid.New(),
		Phone:        phone,
		PasswordHash: "seeded", // placeholder, not used for OTP flow
		IsPremium:    false,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Fatalf("failed to seed user: %v", err)
	}
	log.Printf("Seeded test user with ID %s and phone %s", user.ID, user.Phone)
}
