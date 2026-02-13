package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repositories"
	"errors"

	"gorm.io/gorm"
)

type LoanService interface {
	CreateLoan(customerID, bankID, branchID uint, principal float64) (*models.Loan, error)
	RepayLoan(loanID uint, amount float64) (*models.Loan, error)
	GetLoanDetails(loanID uint) (*models.Loan, float64, error)
}

type loanService struct {
	loanRepo     repositories.LoanRepository
	paymentRepo  repositories.LoanPaymentRepository
	db           *gorm.DB
	customerRepo repositories.CustomerRepository
	bankRepo     repositories.BankRepository
	branchRepo   repositories.BranchRepository
}

func NewLoanService(
	loanRepo repositories.LoanRepository,
	paymentRepo repositories.LoanPaymentRepository,
	customerRepo repositories.CustomerRepository,
	bankRepo repositories.BankRepository,
	branchRepo repositories.BranchRepository,
	db *gorm.DB,
) LoanService {
	return &loanService{
		loanRepo:     loanRepo,
		paymentRepo:  paymentRepo,
		db:           db,
		customerRepo: customerRepo,
		bankRepo:     bankRepo,
		branchRepo:   branchRepo,
	}
}

func (s *loanService) CreateLoan(customerID, bankID, branchID uint, principal float64) (*models.Loan, error) {
	if principal <= 0 {
		return nil, errors.New("principal must be positive")
	}

	if _, err := s.customerRepo.GetByID(customerID); err != nil {
		return nil, errors.New("customer does not exist")
	}

	if _, err := s.bankRepo.GetByID(bankID); err != nil {
		return nil, errors.New("bank does not exist")
	}

	branch, err := s.branchRepo.GetByID(branchID)
	if err != nil {
		return nil, errors.New("branch does not exist")
	}
	if branch.BankID != bankID {
		return nil, errors.New("branch does not belong to bank")
	}

	loan := models.Loan{
		CustomerID:      customerID,
		BankID:          bankID,
		BranchID:        branchID,
		Principal:       principal,
		InterestRate:    12.0, // fixed
		RemainingAmount: principal,
	}

	if err := s.loanRepo.Create(&loan); err != nil {
		return nil, err
	}

	return &loan, nil
}

func (s *loanService) RepayLoan(loanID uint, amount float64) (*models.Loan, error) {
	if amount <= 0 {
		return nil, errors.New("repayment amount must be positive")
	}

	var updatedLoan *models.Loan

	err := s.db.Transaction(func(tx *gorm.DB) error {

		loan, err := s.loanRepo.GetByID(loanID)
		if err != nil {
			return err
		}

		if loan.RemainingAmount < amount {
			return errors.New("repayment exceeds remaining amount")
		}

		newRemaining := loan.RemainingAmount - amount

		if err := s.loanRepo.UpdateRemainingAmount(loanID, newRemaining); err != nil {
			return err
		}

		payment := models.LoanPayment{
			LoanID:     loanID,
			AmountPaid: amount,
		}

		if err := s.paymentRepo.Create(&payment); err != nil {
			return err
		}

		loan.RemainingAmount = newRemaining
		updatedLoan = loan

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedLoan, nil
}

func (s *loanService) GetLoanDetails(loanID uint) (*models.Loan, float64, error) {
	loan, err := s.loanRepo.GetByID(loanID)
	if err != nil {
		return nil, 0, err
	}

	// yearly interest
	interestThisYear := loan.RemainingAmount * (loan.InterestRate / 100)

	return loan, interestThisYear, nil
}
