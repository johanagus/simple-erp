package repository

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type RoleRepository interface {
	FindAll() ([]*domain.Role, error)
	GetRolesByID(RoleID int) ([]string, error)
}
type roleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepositoryImpl{db}
}

func (r *roleRepositoryImpl) FindAll() ([]*domain.Role, error) {
	var roles []*domain.Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *roleRepositoryImpl) GetRolesByID(RoleID int) ([]string, error) {
	var roles []string
	if err := r.db.Table("module_roles").Where("module_roles.role_id = ?", RoleID).Select("module_roles.module_name").Pluck("module_roles.module_name", &roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
