package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAll() (*[]domain.Customer, error)
	FindByID(id int) (*domain.Customer, error)
	Create(customer *domain.Customer) (*domain.Customer, error)
	Update(id int, customer *domain.Customer) (*domain.Customer, error)
	Delete(id uint) error
}

type customerRespositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRespositoryImpl{db: db}
}

func (repo *customerRespositoryImpl) GetAll() (*[]domain.Customer, error) {
	var customers *[]domain.Customer
	result := repo.db.Find(&customers)
	return customers, result.Error
}

func (repo *customerRespositoryImpl) FindByID(id int) (*domain.Customer, error) {
	var customer *domain.Customer
	result := repo.db.Where("id = ?", id).First(&customer)
	return customer, result.Error
}

func (repo *customerRespositoryImpl) Create(customer *domain.Customer) (*domain.Customer, error) {
	result := repo.db.Create(&customer)
	return customer, result.Error
}

func (repo *customerRespositoryImpl) Update(id int, customer *domain.Customer) (*domain.Customer, error) {
	result := repo.db.Where("id = ?", id).Updates(&customer)
	return customer, result.Error
}

func (repo *customerRespositoryImpl) Delete(id uint) error {
	result := repo.db.Delete(&domain.Customer{}, id)
	return result.Error
}
