package service

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
)

type CustomerService interface {
	FindAllCustomer() (*[]domain.Customer, error)
	FindByID(id int) (*domain.Customer, error)
	CreateCustomer(customer *domain.Customer) error
	UpdateCustomer(id int, customer *domain.Customer) error
	DeleteCustomer(id uint)
}

type customerServiceImpl struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerServiceImpl{repo: repo}
}

func (s *customerServiceImpl) FindAllCustomer() (*[]domain.Customer, error) {
	customers, err := s.repo.GetAll()
	return customers, err
}

func (s *customerServiceImpl) FindByID(id int) (*domain.Customer, error) {
	customer, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (s *customerServiceImpl) CreateCustomer(customer *domain.Customer) error {
	_, err := s.repo.Create(customer)
	return err
}

func (s *customerServiceImpl) UpdateCustomer(id int, customer *domain.Customer) error {
	_, err := s.repo.Update(id, customer)
	return err
}

func (s *customerServiceImpl) DeleteCustomer(id uint) {
	s.repo.Delete(id)
}
