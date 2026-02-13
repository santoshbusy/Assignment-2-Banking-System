package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repositories"
	"errors"

	"gorm.io/gorm"
)

type AccountService interface {
	OpenAccount(customerID, bankID, branchID uint) (*models.Account, error)
	Deposit(accountID uint, amount float64) (*models.Account, error)
	Withdraw(accountID uint, amount float64) (*models.Account, error)
	GetAccount(accountID uint) (*models.Account, error)
	Transfer(fromID, toID uint, amount float64) error
	GetTransactions(accountID uint) ([]models.Transaction, error)
}

type accountService struct {
	accountRepo     repositories.AccountRepository
	transactionRepo repositories.TransactionRepository
	db              *gorm.DB
	customerRepo    repositories.CustomerRepository
	branchRepo      repositories.BranchRepository
	bankRepo        repositories.BankRepository
}

func NewAccountService(
	accountRepo repositories.AccountRepository,
	transactionRepo repositories.TransactionRepository,
	customerRepo repositories.CustomerRepository,
	branchRepo repositories.BranchRepository,
	bankRepo repositories.BankRepository,
	db *gorm.DB,
) AccountService {
	return &accountService{
		accountRepo:     accountRepo,
		transactionRepo: transactionRepo,
		customerRepo:    customerRepo,
		branchRepo:      branchRepo,
		bankRepo:        bankRepo,
		db:              db,
	}
}

func (s *accountService) Deposit(accountID uint, amount float64) (*models.Account, error) {
	if amount <= 0 {
		return nil, errors.New("deposit amount must be positive")
	}

	var updatedAccount *models.Account

	err := s.db.Transaction(func(tx *gorm.DB) error {
		account, err := s.accountRepo.GetByID(accountID)
		if err != nil {
			return err
		}

		newBalance := account.Balance + amount
		if err := s.accountRepo.UpdateBalance(accountID, newBalance); err != nil {
			return err
		}
		account.Balance = newBalance

		transaction := models.Transaction{
			AccountID: accountID,
			Type:      "DEPOSIT",
			Amount:    amount,
		}

		if err := s.transactionRepo.Create(&transaction); err != nil {
			return err
		}

		updatedAccount = account
		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

func (s *accountService) Withdraw(accountID uint, amount float64) (*models.Account, error) {
	if amount <= 0 {
		return nil, errors.New("withdraw amount must be positive")
	}

	var updatedAccount *models.Account

	err := s.db.Transaction(func(tx *gorm.DB) error {
		account, err := s.accountRepo.GetByID(accountID)
		if err != nil {
			return err
		}

		if account.Balance < amount {
			return errors.New("insufficient balance")
		}

		newBalance := account.Balance - amount

		if err := s.accountRepo.UpdateBalance(accountID, newBalance); err != nil {
			return err
		}

		transaction := models.Transaction{
			AccountID: accountID,
			Type:      "WITHDRAW",
			Amount:    amount,
		}

		if err := s.transactionRepo.Create(&transaction); err != nil {
			return err
		}

		account.Balance = newBalance
		updatedAccount = account

		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}

func (s *accountService) GetAccount(accountID uint) (*models.Account, error) {
	return s.accountRepo.GetByID(accountID)
}

func (s *accountService) Transfer(fromID, toID uint, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {

		fromAccount, err := s.accountRepo.GetByID(fromID)
		if err != nil {
			return err
		}

		toAccount, err := s.accountRepo.GetByID(toID)
		if err != nil {
			return err
		}

		if fromAccount.Balance < amount {
			return errors.New("insufficient balance")
		}

		// Deduct from sender
		if err := s.accountRepo.UpdateBalance(fromID, fromAccount.Balance-amount); err != nil {
			return err
		}

		// Add to receiver
		if err := s.accountRepo.UpdateBalance(toID, toAccount.Balance+amount); err != nil {
			return err
		}

		// Create transaction records
		if err := s.transactionRepo.Create(&models.Transaction{
			AccountID: fromID,
			Type:      "TRANSFER_OUT",
			Amount:    amount,
		}); err != nil {
			return err
		}

		if err := s.transactionRepo.Create(&models.Transaction{
			AccountID: toID,
			Type:      "TRANSFER_IN",
			Amount:    amount,
		}); err != nil {
			return err
		}

		return nil
	})
}

func (s *accountService) GetTransactions(accountID uint) ([]models.Transaction, error) {
	return s.transactionRepo.GetByAccountID(accountID)
}

func (s *accountService) OpenAccount(customerID, bankID, branchID uint) (*models.Account, error) {

	// Validate customer exists
	_, err := s.customerRepo.GetByID(customerID)
	if err != nil {
		return nil, errors.New("customer does not exist")
	}

	// Validate bank exists
	if _, err := s.bankRepo.GetByID(bankID); err != nil {
		return nil, errors.New("bank does not exist")
	}

	// Validate branch exists and belongs to the bank
	branch, err := s.branchRepo.GetByID(branchID)
	if err != nil {
		return nil, errors.New("branch does not exist")
	}
	if branch.BankID != bankID {
		return nil, errors.New("branch does not belong to bank")
	}

	account := models.Account{
		CustomerID: customerID,
		BankID:     bankID,
		BranchID:   branchID,
		Balance:    0,
	}

	if err := s.accountRepo.Create(&account); err != nil {
		return nil, err
	}

	return &account, nil
}
