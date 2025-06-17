package service

import (
	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
)

type InventoriService interface {
	FindByID(id int) (*domain.Inventory, error)
	FindByWHID(whID int) (*[]domain.Inventory, error)
	SaveInventori(inventori *domain.Inventory) (bool, error)
	UpdateInventori(inventori *domain.Inventory) (*domain.Inventory, error)
}

type inventoriServiceImpl struct {
	invRepo repository.InventoryRepository
	whRepo  repository.WarehouseRepository
}

func NewInventoriService(invRepo repository.InventoryRepository, whRepo repository.WarehouseRepository) InventoriService {
	return &inventoriServiceImpl{invRepo: invRepo, whRepo: whRepo}
}

func (s *inventoriServiceImpl) FindByID(id int) (*domain.Inventory, error) {
	inventori, err := s.invRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return inventori, nil
}

func (s *inventoriServiceImpl) FindByWHID(whID int) (*[]domain.Inventory, error) {
	inventories, err := s.invRepo.FindByWHID(whID)
	if err != nil {
		return nil, err
	}
	return inventories, nil
}

func (s *inventoriServiceImpl) SaveInventori(inventori *domain.Inventory) (bool, error) {
	// Cek warehouse ID pada inventori
	whID := inventori.WarehouseID
	wh, err := s.whRepo.FindByID(whID)
	if err != nil || wh == nil {
		return false, err
	}

	success, err := s.invRepo.SaveInventori(inventori)

	if err != nil {
		return false, err
	}
	return success, nil
}

func (s *inventoriServiceImpl) UpdateInventori(inventori *domain.Inventory) (*domain.Inventory, error) {
	updatedInventori, err := s.invRepo.UpdateInventori(inventori)
	if err != nil {
		return nil, err
	}
	return updatedInventori, nil
}
