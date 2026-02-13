package repositories

import (
	"banking-system/internal/models"

	"gorm.io/gorm"
)

type LoanRepository interface {
	Create(loan *models.Loan) error
	GetByID(id uint) (*models.Loan, error)
	UpdateRemainingAmount(loanID uint, amount float64) error
}

type loanRepository struct {
	db *gorm.DB
}

func NewLoanRepository(db *gorm.DB) LoanRepository {
	return &loanRepository{db: db}
}

func (r *loanRepository) Create(loan *models.Loan) error {
	return r.db.Create(loan).Error
}

func (r *loanRepository) GetByID(id uint) (*models.Loan, error) {
	var loan models.Loan
	err := r.db.First(&loan, "loan_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &loan, nil
}

func (r *loanRepository) UpdateRemainingAmount(loanID uint, amount float64) error {
	return r.db.Model(&models.Loan{}).
		Where("loan_id = ?", loanID).
		Update("remaining_amount", amount).Error
}
