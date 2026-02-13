package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repositories"
)

type CustomerService interface {
	CreateCustomer(name, email string) (*models.Customer, error)
	GetCustomer(id uint) (*models.Customer, error)
}

type customerService struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) CustomerService {
	return &customerService{customerRepo: customerRepo}
}

func (s *customerService) CreateCustomer(name, email string) (*models.Customer, error) {
	customer := models.Customer{
		Name:  name,
		Email: email,
	}
	if err := s.customerRepo.Create(&customer); err != nil {
		return nil, err
	}
	return &customer, nil
}

func (s *customerService) GetCustomer(id uint) (*models.Customer, error) {
	return s.customerRepo.GetByID(id)
}
