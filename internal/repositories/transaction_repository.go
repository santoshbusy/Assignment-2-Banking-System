package repositories

import (
	"banking-system/internal/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *models.Transaction) error
	GetByAccountID(accountID uint) ([]models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(tx *models.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *transactionRepository) GetByAccountID(accountID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.
		Where("account_id = ?", accountID).
		Order("created_at desc").
		Find(&transactions).Error

	return transactions, err
}
