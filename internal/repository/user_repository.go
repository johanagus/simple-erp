package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindByID(id int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	SaveUser(user *domain.User) error
	UpdateUser(id int, user *domain.User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (repo *userRepositoryImpl) FindAll() ([]domain.User, error) {
	var users []domain.User
	result := repo.db.Find(&users)
	return users, result.Error
}

func (repo *userRepositoryImpl) FindByID(id int) (domain.User, error) {
	var user domain.User
	result := repo.db.Where("id = ?", id).First(&user)
	return user, result.Error
}

func (repo *userRepositoryImpl) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	result := repo.db.Where("email = ?", email).Find(&user)
	return user, result.Error
}

func (repo *userRepositoryImpl) SaveUser(user *domain.User) error {
	result := repo.db.Create(user)
	return result.Error
}

func (repo *userRepositoryImpl) UpdateUser(id int, user *domain.User) error {
	result := repo.db.Where("id = ?", id).Updates(user)
	return result.Error
}
