package db

import (
	"banking-system/internal/models"
	"log"

	"gorm.io/gorm"
)

// AutoMigrate runs GORM schema migrations for all domain models.
// Tables are created in dependency order so parents exist before children.
func AutoMigrate(database *gorm.DB) {
	if database == nil {
		log.Println("skipping auto-migration: database connection is nil")
		return
	}

	// Migrate in dependency order: parents before children
	batches := []interface{}{
		&models.Bank{},
		&models.Customer{},
		&models.Branch{},
		&models.Account{},
		&models.Transaction{},
		&models.Loan{},
		&models.LoanPayment{},
	}

	for _, model := range batches {
		if err := database.AutoMigrate(model); err != nil {
			log.Fatalf("auto-migration failed: %v", err)
		}
	}

	log.Println("database schema is up to date")
}
