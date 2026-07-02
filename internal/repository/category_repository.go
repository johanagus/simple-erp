package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAll() ([]domain.Category, error)
	GetByID(id int) (domain.Category, error)
	Create(category domain.Category) (domain.Category, error)
	Update(category domain.Category) (domain.Category, error)
	Delete(id int) error
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}

func (repo *categoryRepositoryImpl) GetAll() ([]domain.Category, error) {
	var categories []domain.Category
	result := repo.db.Find(&categories)
	return categories, result.Error
}

func (repo *categoryRepositoryImpl) GetByID(id int) (domain.Category, error) {
	var category domain.Category
	result := repo.db.Where("id = ?", id).First(&category)
	return category, result.Error
}

func (repo *categoryRepositoryImpl) Create(category domain.Category) (domain.Category, error) {
	result := repo.db.Create(&category)
	return category, result.Error
}

func (repo *categoryRepositoryImpl) Update(category domain.Category) (domain.Category, error) {
	result := repo.db.Save(&category)
	return category, result.Error
}

func (repo *categoryRepositoryImpl) Delete(id int) error {
	result := repo.db.Delete(&domain.Category{}, id)
	return result.Error
}
