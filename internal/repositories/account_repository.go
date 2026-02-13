package repositories

import (
	"banking-system/internal/models"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(account *models.Account) error
	GetByID(id uint) (*models.Account, error)
	UpdateBalance(accountID uint, newBalance float64) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) GetByID(id uint) (*models.Account, error) {
	var account models.Account
	err := r.db.First(&account, "account_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) UpdateBalance(accountID uint, newBalance float64) error {
	return r.db.Model(&models.Account{}).
		Where("account_id = ?", accountID).
		Update("balance", newBalance).Error
}
