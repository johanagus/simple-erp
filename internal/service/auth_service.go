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
	repo repository.AuthRespository
}

func NewAuthService(repo repository.AuthRespository) AuthService {
	return &authServiceImpl{repo}
}

func (s *authServiceImpl) Authenticate(email, password string) (*domain.User, error) {
	user, err := s.repo.FindByEmail(email)

	if err != nil {
		return nil, errors.New("email tidak ditemukan")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("password salah")
	}

	return user, nil

}
