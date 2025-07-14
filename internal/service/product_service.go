package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/johanagus/simple-erp/internal/domain"
	"github.com/johanagus/simple-erp/internal/dto"
	"github.com/johanagus/simple-erp/internal/repository"
)

type ProductService interface {
	FindAll() ([]domain.Product, error)
	FindByID(id int) (*domain.Product, error)
	FindBySKU(sku string) (*domain.Product, error)
	FindByBarcode(barcode string) (*domain.Product, error)
	Search(query string) ([]domain.Product, error)
	SaveProduct(p *dto.ProductRequest) (int, error)
	UpdateProduct(id int, product *domain.Product) (*domain.Product, error)
}

type productServiceImpl struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewProductService(productRepo repository.ProductRepository, categoryRepo repository.CategoryRepository) ProductService {
	return &productServiceImpl{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *productServiceImpl) FindAll() ([]domain.Product, error) {
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productServiceImpl) FindByID(id int) (*domain.Product, error) {
	return s.productRepo.FindByID(id)
}

func (s *productServiceImpl) FindBySKU(sku string) (*domain.Product, error) {
	return s.productRepo.FindBySKU(sku)
}

func (s *productServiceImpl) FindByBarcode(barcode string) (*domain.Product, error) {
	return s.productRepo.FindByBarcode(barcode)
}

func (s *productServiceImpl) Search(query string) ([]domain.Product, error) {
	return s.productRepo.Search(query)
}

func (s *productServiceImpl) SaveProduct(p *dto.ProductRequest) (int, error) {
	now := time.Now()
	year := now.Year()
	month := int(now.Month())

	// Find Category
	category, err := s.categoryRepo.FindByName(p.Category)
	if err != nil {
		return 0, err
	}

	if category == nil {
		return 0, errors.New("category not found")
	}

	var urutan int

	LastSKUID := s.productRepo.FindLastSKUByCategoryAndPeriod(year, month, category.ID)

	if LastSKUID != "" && len(LastSKUID) >= 4 {
		fmt.Printf("jalan")
		// Ambil 4 digit terakhir dari SKU terakhir
		fmt.Sscanf(LastSKUID[len(LastSKUID)-4:], "%04d", &urutan)
	}

	urutan++

	fmt.Println(urutan)

	sku := fmt.Sprintf("%02d%02d%d%04d", year%100, month, category.ID, urutan)

	product := &domain.Product{
		SKU:         sku,
		Barcode:     p.Barcode,
		Name:        p.Name,
		CategoryID:  category.ID,
		Type:        p.Type,
		Brand:       p.Brand,
		PriceLevel1: p.Price,
		IsSN:        p.IsSN,
	}

	// Pastikan SKU unik
	res, err := s.productRepo.FindBySKU(sku)
	if err == nil && res != nil {
		return 0, errors.New("sku already exists")
	}

	return s.productRepo.SaveProduct(product)
}

func (s *productServiceImpl) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	return s.productRepo.UpdateProduct(id, product)
}
