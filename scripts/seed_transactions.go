package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"qotera-backend/internal/domain"
	"qotera-backend/pkg/config"
	"qotera-backend/pkg/database"

	"github.com/google/uuid"
)

// Seed specific data testing for the user: "00000000-0000-0000-0000-000000000001"
var testUserID = uuid.MustParse("00000000-0000-0000-0000-000000000001")

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Seeding Test User and Past 7 Days of Transactions...")

	// 1. Ensure the Test User Exists
	testUser := domain.User{
		ID:    testUserID,
		Phone: "+251912345678",
	}

	// upsert user
	if err := db.Where("id = ?", testUserID).FirstOrCreate(&testUser).Error; err != nil {
		log.Fatalf("Failed to seed test user: %v", err)
	}

	// Clear old transactions for this user to avoid messy overlaps during testing
	db.Where("user_id = ?", testUserID).Delete(&domain.Transaction{})

	now := time.Now()
	var seededTransactions []domain.Transaction

	for i := 0; i < 7; i++ {
		// Generate 1-4 random transactions per day for the last week
		numTxPerDay := rand.Intn(4) + 1
		txDate := now.AddDate(0, 0, -i) // Go back 'i' days

		for j := 0; j < numTxPerDay; j++ {
			isIncome := rand.Intn(2) == 0
			isCBE := rand.Intn(2) == 0

			amount := float64(rand.Intn(4900) + 100) // 100 to 5000 ETB
			var fee *float64

			txType := domain.TypeExpense
			if isIncome {
				txType = domain.TypeIncome
			} else {
				// Add a small randomized fee for expenses
				f := float64(rand.Intn(40) + 5)
				fee = &f
			}

			source := domain.SourceTelebirr
			if isCBE {
				source = domain.SourceCBE
			}

			// Generate a random generic reference number format to ensure uniqueness
			refNum := fmt.Sprintf("REF-%s-%d%d", source, txDate.Unix(), j)

			tx := domain.Transaction{
				UserID:          testUserID,
				Amount:          amount,
				Fee:             fee,
				Type:            txType,
				Source:          source,
				ReferenceNumber: &refNum,
				TransactionDate: txDate,
				RawMessage:      "Seeded DB Record for UI Graph Visualization Testing",
			}

			seededTransactions = append(seededTransactions, tx)
		}
	}

	// Insert into Postgres Using GORM
	err = db.WithContext(context.Background()).Create(&seededTransactions).Error
	if err != nil {
		log.Fatalf("Failed to insert seeded transactions: %v", err)
	}

	log.Printf("Successfully Seeded %d transactions for User %s over the past 7 days!\n", len(seededTransactions), testUserID.String())
}
