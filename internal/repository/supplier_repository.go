package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	FindAll() (*[]domain.Supplier, error)
	FindByID(id int) (*domain.Supplier, error)
	FindByName(name string) (*[]domain.Supplier, error)
	SaveSupplier(supplier *domain.Supplier) (bool, error)
	UpdateSupplier(id int, supplier *domain.Supplier) (*domain.Supplier, error)
	DeleteSupplier(id int) (bool, error)
}

type supplierRepositoryImpl struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepositoryImpl{db: db}
}

func (repo *supplierRepositoryImpl) FindAll() (*[]domain.Supplier, error) {
	var suppliers []domain.Supplier
	result := repo.db.Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}

	return &suppliers, nil
}

func (repo *supplierRepositoryImpl) FindByID(id int) (*domain.Supplier, error) {
	var supplier domain.Supplier
	result := repo.db.Where("id = ?", id).First(&supplier)
	if result.Error != nil {
		return nil, result.Error
	}
	return &supplier, nil
}

func (repo *supplierRepositoryImpl) FindByName(name string) (*[]domain.Supplier, error) {
	var suppliers []domain.Supplier
	result := repo.db.Where("name LIKE ?", "%"+name+"%").Find(&suppliers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &suppliers, nil
}

func (repo *supplierRepositoryImpl) SaveSupplier(supplier *domain.Supplier) (bool, error) {
	result := repo.db.Create(supplier)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (repo *supplierRepositoryImpl) UpdateSupplier(id int, supplier *domain.Supplier) (*domain.Supplier, error) {
	result := repo.db.Where("id = ?", id).Updates(supplier)
	if result.Error != nil {
		return nil, result.Error
	}
	return supplier, nil
}

func (repo *supplierRepositoryImpl) DeleteSupplier(id int) (bool, error) {
	var supplier domain.Supplier
	result := repo.db.Where("id = ?", id).First(&supplier)
	if result.Error != nil {
		return false, result.Error
	}

	if err := repo.db.Delete(&supplier).Error; err != nil {
		return false, err
	}
	return true, nil
}
