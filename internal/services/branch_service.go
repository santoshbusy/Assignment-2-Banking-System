package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repositories"
)

type BranchService interface {
	CreateBranch(bankID uint, name, address string) (*models.Branch, error)
	GetBranches(bankID uint) ([]models.Branch, error)
}

type branchService struct {
	branchRepo repositories.BranchRepository
}

func NewBranchService(branchRepo repositories.BranchRepository) BranchService {
	return &branchService{branchRepo: branchRepo}
}

func (s *branchService) CreateBranch(bankID uint, name, address string) (*models.Branch, error) {
	branch := models.Branch{
		BankID:  bankID,
		Name:    name,
		Address: address,
	}
	if err := s.branchRepo.Create(&branch); err != nil {
		return nil, err
	}
	return &branch, nil
}

func (s *branchService) GetBranches(bankID uint) ([]models.Branch, error) {
	return s.branchRepo.GetByBank(bankID)
}
