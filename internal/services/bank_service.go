package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repositories"
)

type BankService interface {
	CreateBank(name string) (*models.Bank, error)
	GetBanks() ([]models.Bank, error)
}

type bankService struct {
	bankRepo repositories.BankRepository
}

func NewBankService(bankRepo repositories.BankRepository) BankService {
	return &bankService{bankRepo: bankRepo}
}

func (s *bankService) CreateBank(name string) (*models.Bank, error) {
	bank := models.Bank{
		Name: name,
	}
	if err := s.bankRepo.Create(&bank); err != nil {
		return nil, err
	}
	return &bank, nil
}

func (s *bankService) GetBanks() ([]models.Bank, error) {
	return s.bankRepo.GetAll()
}
