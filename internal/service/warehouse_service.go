// buatkan satu service untuk warehouse
package service

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
)

type WarehouseService interface {
	FindAllWarehouses() ([]domain.Warehouse, error)
	FindWarehouseByID(id int) (*domain.Warehouse, error)
	CreateWarehouse(warehouse *domain.Warehouse) (int, error)
	UpdateWarehouse(warehouse *domain.Warehouse) (*domain.Warehouse, error)
}

type warehouseServiceImpl struct {
	repo repository.WarehouseRepository
}

func NewWarehouseService(repo repository.WarehouseRepository) WarehouseService {
	return &warehouseServiceImpl{repo: repo}
}

func (s *warehouseServiceImpl) FindAllWarehouses() ([]domain.Warehouse, error) {
	return s.repo.FindAll()
}

func (s *warehouseServiceImpl) FindWarehouseByID(id int) (*domain.Warehouse, error) {
	return s.repo.FindByID(id)
}

func (s *warehouseServiceImpl) CreateWarehouse(warehouse *domain.Warehouse) (int, error) {
	return s.repo.SaveWarehouse(warehouse)
}

func (s *warehouseServiceImpl) UpdateWarehouse(warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	return s.repo.UpdateWarehouse(warehouse)
}
