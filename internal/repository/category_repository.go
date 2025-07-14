package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll() []domain.Category
	FindByID(id int) *domain.Category
	FindByName(name string) (*domain.Category, error)
}

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepositoryImpl{DB: db}
}

func (repo CategoryRepositoryImpl) FindAll() []domain.Category {
	var categories []domain.Category
	repo.DB.Find(&categories)
	return categories
}

func (repo CategoryRepositoryImpl) FindByID(id int) *domain.Category {
	var category domain.Category
	repo.DB.Find(&category, id)
	return &category
}

func (repo CategoryRepositoryImpl) FindByName(name string) (*domain.Category, error) {
	var category domain.Category
	result := repo.DB.First(&category, "name = ?", name)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Tidak ditemukan, return nil
		}
		return nil, result.Error // Error lain
	}
	return &category, nil
}
