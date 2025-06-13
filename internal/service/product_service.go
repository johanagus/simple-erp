package service

import (
	"errors"

	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/repository"
)

type ProductService interface {
	FindAll() ([]domain.Product, error)
	FindByID(id int) (*domain.Product, error)
	FindBySKU(sku string) (*domain.Product, error)
	FindByBarcode(barcode string) (*domain.Product, error)
	Search(query string) ([]domain.Product, error)
	SaveProduct(product *domain.Product) (int, error)
	UpdateProduct(product *domain.Product) (*domain.Product, error)
}

type productServiceImpl struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}

func (s *productServiceImpl) FindAll() ([]domain.Product, error) {
	return s.repo.FindAll()
}

func (s *productServiceImpl) FindByID(id int) (*domain.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productServiceImpl) FindBySKU(sku string) (*domain.Product, error) {
	return s.repo.FindBySKU(sku)
}

func (s *productServiceImpl) FindByBarcode(barcode string) (*domain.Product, error) {
	return s.repo.FindByBarcode(barcode)
}

func (s *productServiceImpl) Search(query string) ([]domain.Product, error) {
	return s.repo.Search(query)
}

func (s *productServiceImpl) SaveProduct(product *domain.Product) (int, error) {
	sku := product.SKU
	res, err := s.repo.FindBySKU(sku)
	if err == nil && res != nil {
		return 0, errors.New("sku already exists")
	}
	return s.repo.SaveProduct(product)
}

func (s *productServiceImpl) UpdateProduct(product *domain.Product) (*domain.Product, error) {
	return s.repo.UpdateProduct(product)
}
