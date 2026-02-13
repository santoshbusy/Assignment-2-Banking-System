package repositories

import (
	"banking-system/internal/models"

	"gorm.io/gorm"
)

type LoanPaymentRepository interface {
	Create(payment *models.LoanPayment) error
	GetByLoanID(loanID uint) ([]models.LoanPayment, error)
}

type loanPaymentRepository struct {
	db *gorm.DB
}

func NewLoanPaymentRepository(db *gorm.DB) LoanPaymentRepository {
	return &loanPaymentRepository{db: db}
}

func (r *loanPaymentRepository) Create(payment *models.LoanPayment) error {
	return r.db.Create(payment).Error
}

func (r *loanPaymentRepository) GetByLoanID(loanID uint) ([]models.LoanPayment, error) {
	var payments []models.LoanPayment
	err := r.db.
		Where("loan_id = ?", loanID).
		Order("created_at desc").
		Find(&payments).Error

	return payments, err
}
