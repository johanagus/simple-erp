package service

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
)

type SupplierService interface {
	FindAll() (*[]domain.Supplier, error)
	FindByID(id int) (*domain.Supplier, error)
	FindByName(name string) (*[]domain.Supplier, error)
	SaveSupplier(supplier *domain.Supplier) (bool, error)
	UpdateSupplier(id int, supplier *domain.Supplier) (*domain.Supplier, error)
	DeleteSupplier(id int) (bool, error)
}
type supplierServiceImpl struct {
	repo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierServiceImpl{repo: repo}
}

func (s *supplierServiceImpl) FindAll() (*[]domain.Supplier, error) {
	suppliers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if suppliers == nil || len(*suppliers) == 0 {
		return nil, nil // Return nil if no suppliers found
	}

	return suppliers, nil
}

func (s *supplierServiceImpl) FindByID(id int) (*domain.Supplier, error) {
	supplier, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

func (s *supplierServiceImpl) FindByName(name string) (*[]domain.Supplier, error) {
	suppliers, err := s.repo.FindByName(name)
	if err != nil {
		return nil, err
	}
	return suppliers, nil
}

func (s *supplierServiceImpl) SaveSupplier(supplier *domain.Supplier) (bool, error) {
	success, err := s.repo.SaveSupplier(supplier)
	if err != nil {
		return false, err
	}
	return success, nil
}

func (s *supplierServiceImpl) UpdateSupplier(id int, supplier *domain.Supplier) (*domain.Supplier, error) {
	updatedSupplier, err := s.repo.UpdateSupplier(id, supplier)
	if err != nil {
		return nil, err
	}
	return updatedSupplier, nil
}

func (s *supplierServiceImpl) DeleteSupplier(id int) (bool, error) {
	success, err := s.repo.DeleteSupplier(id)
	if err != nil {
		return false, err
	}
	return success, nil
}
