package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type WarehouseRepository interface {
	FindAll() ([]domain.Warehouse, error)
	FindByID(id int) (*domain.Warehouse, error)
	SaveWarehouse(warehouse *domain.Warehouse) (int, error)
	UpdateWarehouse(warehouse *domain.Warehouse) (*domain.Warehouse, error)
}

type warehouseRepositoryImpl struct {
	DB *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) WarehouseRepository {
	return &warehouseRepositoryImpl{DB: db}
}

func (repo *warehouseRepositoryImpl) FindAll() ([]domain.Warehouse, error) {
	var warehouses []domain.Warehouse
	result := repo.DB.Find(&warehouses)
	return warehouses, result.Error
}

func (repo *warehouseRepositoryImpl) FindByID(id int) (*domain.Warehouse, error) {
	var warehouse *domain.Warehouse
	result := repo.DB.Where("id = ?", id).First(&warehouse)
	return warehouse, result.Error
}

func (repo *warehouseRepositoryImpl) SaveWarehouse(warehouse *domain.Warehouse) (int, error) {
	result := repo.DB.Create(warehouse)
	id := result.RowsAffected
	return int(id), result.Error
}

func (repo *warehouseRepositoryImpl) UpdateWarehouse(warehouse *domain.Warehouse) (*domain.Warehouse, error) {
	result := repo.DB.Save(warehouse)
	return warehouse, result.Error
}
