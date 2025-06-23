package service

import (
	"errors"
	"fmt"

	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
	"github.com/johanagus/simple-erp/pkg/utils"
)

type UserService interface {
	FindAll() ([]domain.User, error)
	FindByID(id int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	SaveUser(user *domain.User) error
	UpdateUser(id int, user *domain.User) error
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) FindAll() ([]domain.User, error) {

	users, err := s.repo.FindAll()
	if err != nil {
		return []domain.User{}, errors.New("user tidak di temukan")
	}

	return users, nil

}

func (s *userServiceImpl) FindByID(id int) (domain.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return domain.User{}, errors.New("user tidak di temukan")
	}

	return user, nil
}

func (s *userServiceImpl) FindByEmail(email string) (domain.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return domain.User{}, errors.New("user tidak di temukan")
	}

	return user, nil
}

func (s *userServiceImpl) SaveUser(user *domain.User) error {
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return errors.New("gagal hash password")
	}
	fmt.Println(user)
	user.Password = hashPassword
	fmt.Println(user)

	err = s.repo.SaveUser(user)
	if err != nil {
		return errors.New("gagal menyimpan user")
	}

	return nil

}

func (s *userServiceImpl) UpdateUser(id int, user *domain.User) error {
	err := s.repo.UpdateUser(id, user)
	if err != nil {
		return errors.New("gagal memperbarui user")
	}
	return nil
}
