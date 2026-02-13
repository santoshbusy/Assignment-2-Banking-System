package repositories

import (
	"banking-system/internal/models"

	"gorm.io/gorm"
)

type BankRepository interface {
	Create(bank *models.Bank) error
	GetAll() ([]models.Bank, error)
	GetByID(id uint) (*models.Bank, error)
}

type bankRepository struct {
	db *gorm.DB
}

func NewBankRepository(db *gorm.DB) BankRepository {
	return &bankRepository{db: db}
}

func (r *bankRepository) Create(bank *models.Bank) error {
	return r.db.Create(bank).Error
}

func (r *bankRepository) GetAll() ([]models.Bank, error) {
	var banks []models.Bank
	err := r.db.Find(&banks).Error
	return banks, err
}

func (r *bankRepository) GetByID(id uint) (*models.Bank, error) {
	var bank models.Bank
	if err := r.db.First(&bank, "bank_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &bank, nil
}
