package service

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
)

type StoreService interface {
	GetAll() (*[]domain.Store, error)
	FindByID(id int) (*domain.Store, error)
	Create(store *domain.Store) (*domain.Store, error)
	Update(id int, store *domain.Store) (*domain.Store, error)
	Delete(id int) error
}

type storeServiceImpl struct {
	repo repository.StoreRepository
}

func NewStoreService(repo repository.StoreRepository) StoreService {
	return &storeServiceImpl{repo: repo}
}

func (s *storeServiceImpl) GetAll() (*[]domain.Store, error) {
	return s.repo.FindAll(), nil
}

func (s *storeServiceImpl) FindByID(id int) (*domain.Store, error) {
	return s.repo.FindByID(id), nil
}

func (s *storeServiceImpl) Create(store *domain.Store) (*domain.Store, error) {
	return s.repo.CreateStore(store)
}

func (s *storeServiceImpl) Update(id int, store *domain.Store) (*domain.Store, error) {
	return s.repo.UpdateStore(id, store)
}

func (s *storeServiceImpl) Delete(id int) error {
	_, err := s.repo.DeleteStore(int(id))
	return err
}
