package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	FindByID(id int) (*domain.Inventory, error)
	FindByWHID(whID int) (*[]domain.Inventory, error)
	SaveInventori(inventori *domain.Inventory) (bool, error)
	UpdateInventori(inventori *domain.Inventory) (*domain.Inventory, error)
}

type inventoriRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoriRepository(db *gorm.DB) InventoryRepository {
	return &inventoriRepositoryImpl{db: db}
}

func (repo *inventoriRepositoryImpl) FindByID(id int) (*domain.Inventory, error) {
	var inventori *domain.Inventory
	result := repo.db.Where("id = ?", id).First(&inventori)
	return inventori, result.Error
}

func (repo *inventoriRepositoryImpl) FindByWHID(whID int) (*[]domain.Inventory, error) {
	var inventories *[]domain.Inventory
	result := repo.db.Where("warehouse_id", whID).Find(&inventories)
	return inventories, result.Error
}

func (repo *inventoriRepositoryImpl) SaveInventori(inventori *domain.Inventory) (bool, error) {
	result := repo.db.Save(inventori)
	return result.Error == nil, result.Error
}

func (repo *inventoriRepositoryImpl) UpdateInventori(inventori *domain.Inventory) (*domain.Inventory, error) {
	result := repo.db.Save(inventori)
	return inventori, result.Error
}
