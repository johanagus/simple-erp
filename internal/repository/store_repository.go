package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type StoreRepository interface {
	FindAll() *[]domain.Store
	FindByID(id int) *domain.Store
	CreateStore(store *domain.Store) (*domain.Store, error)
	UpdateStore(id int, store *domain.Store) (*domain.Store, error)
	DeleteStore(id int) (bool, error)
}

type storeRepositoryImpl struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepositoryImpl{db: db}
}

func (repo *storeRepositoryImpl) FindAll() *[]domain.Store {
	var stores []domain.Store
	repo.db.Find(&stores)
	return &stores
}

func (repo *storeRepositoryImpl) FindByID(id int) *domain.Store {
	var store domain.Store
	repo.db.Where("id = ?", id).First(&store)
	return &store
}

func (repo *storeRepositoryImpl) CreateStore(store *domain.Store) (*domain.Store, error) {
	result := repo.db.Create(store)
	if result.Error != nil {
		return nil, result.Error
	}
	return store, nil
}

func (repo *storeRepositoryImpl) UpdateStore(id int, store *domain.Store) (*domain.Store, error) {
	result := repo.db.Where("id = ?", id).Updates(store)
	if result.Error != nil {
		return nil, result.Error
	}
	return store, nil
}

func (repo *storeRepositoryImpl) DeleteStore(id int) (bool, error) {
	var store domain.Store
	result := repo.db.Where("id = ?", id).First(&store)
	if result.Error != nil {
		return false, result.Error
	}
	if err := repo.db.Delete(&store).Error; err != nil {
		return false, err
	}
	return true, nil
}
