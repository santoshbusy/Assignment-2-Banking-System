package repositories

import (
	"banking-system/internal/models"

	"gorm.io/gorm"
)

type BranchRepository interface {
	Create(branch *models.Branch) error
	GetByBank(bankID uint) ([]models.Branch, error)
	GetByID(id uint) (*models.Branch, error)
}

type branchRepository struct {
	db *gorm.DB
}

func NewBranchRepository(db *gorm.DB) BranchRepository {
	return &branchRepository{db: db}
}

func (r *branchRepository) Create(branch *models.Branch) error {
	return r.db.Create(branch).Error
}

func (r *branchRepository) GetByBank(bankID uint) ([]models.Branch, error) {
	var branches []models.Branch
	err := r.db.Preload("Bank").Where("bank_id = ?", bankID).Find(&branches).Error
	return branches, err
}

func (r *branchRepository) GetByID(id uint) (*models.Branch, error) {
	var branch models.Branch
	if err := r.db.First(&branch, "branch_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &branch, nil
}
