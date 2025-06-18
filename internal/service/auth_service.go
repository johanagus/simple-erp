package service

import (
	"errors"

	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
	"github.com/johanagus/simple-erp/pkg/utils"
)

type AuthService interface {
	Authenticate(username, password string) (*domain.User, error)
}

type authServiceImpl struct {
	authRepo repository.AuthRespository
	roleRepo repository.RoleRepository
}

func NewAuthService(repo repository.AuthRespository, roleRepo repository.RoleRepository) AuthService {
	return &authServiceImpl{authRepo: repo, roleRepo: roleRepo}
}

func (s *authServiceImpl) Authenticate(email, password string) (*domain.User, error) {
	user, err := s.authRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("email tidak ditemukan")
	}

	roles, err := s.roleRepo.GetRolesByID(user.RoleID)
	if err != nil {
		return nil, errors.New("role tidak ditemukan")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("password salah")
	}

	user.Roles = roles
	return user, nil

}
