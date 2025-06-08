package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type AuthRespository interface {
	FindByEmail(username string) (*domain.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRespository {
	return &authRepository{db: db}
}

func (r *authRepository) FindByEmail(email string) (*domain.User, error) {

	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
