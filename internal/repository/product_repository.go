package repository

import (
	"strconv"

	"github.com/johanagus/simple-erp/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]domain.Product, error)
	FindByID(id int) (*domain.Product, error)
	FindByBarcode(barcode string) (*domain.Product, error)
	FindBySKU(sku string) (*domain.Product, error)
	Search(query string) ([]domain.Product, error)
	SaveProduct(product *domain.Product) (int, error)
	UpdateProduct(id int, product *domain.Product) (*domain.Product, error)
	FindLastSKUByCategoryAndPeriod(year, month int, categoryID int) string
}

type productRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (repo *productRepositoryImpl) FindAll() ([]domain.Product, error) {
	var products []domain.Product
	result := repo.DB.Select("*").Find(&products).InnerJoins("category")
	return products, result.Error
}

func (repo *productRepositoryImpl) FindByID(id int) (*domain.Product, error) {
	var product *domain.Product
	result := repo.DB.Where("id = ?", id).First(&product)
	return product, result.Error
}

func (repo *productRepositoryImpl) FindByBarcode(barcode string) (*domain.Product, error) {
	var product *domain.Product
	result := repo.DB.Where("barcode = ?", barcode).First(&product)
	return product, result.Error
}

func (repo *productRepositoryImpl) FindBySKU(sku string) (*domain.Product, error) {
	var product *domain.Product
	result := repo.DB.Where("sku = ?", sku).First(&product)
	return product, result.Error
}

func (repo *productRepositoryImpl) Search(query string) ([]domain.Product, error) {
	var products []domain.Product
	result := repo.DB.Where("name LIKE ?", "%"+query+"%").Find(&products)
	return products, result.Error
}

func (repo *productRepositoryImpl) SaveProduct(product *domain.Product) (int, error) {
	result := repo.DB.Create(product)
	id := result.RowsAffected
	return int(id), result.Error
}

func (repo *productRepositoryImpl) UpdateProduct(id int, product *domain.Product) (*domain.Product, error) {
	result := repo.DB.Where("id = ?", id).Updates(product)
	return product, result.Error
}

func (repo *productRepositoryImpl) FindLastSKUByCategoryAndPeriod(year, month int, categoryID int) string {
	var skuID int
	repo.DB.Raw("SELECT sku FROM products WHERE category_id = ? AND MONTH(created_at) = ? AND YEAR(created_at) = ? ORDER BY id DESC LIMIT 1", categoryID, month, year).Scan(&skuID)
	return strconv.Itoa(skuID)
}
