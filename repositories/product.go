package repositories

import (
	"log"

	"gorm.io/gorm"

	"github.com/hoanggggg5/shopproduct/models"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	Migrate() error
	GetProduct(int64, ...interface{}) (*models.Product, error)
	GetProducts(...interface{}) ([]*models.Product, error)
	CreateProduct(*models.Product) (*models.Product, error)
	UpdateProduct(*models.Product) (*models.Product, error)
	DeleteProduct(id int64) error
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{
		DB: db,
	}
}

func (p productRepository) Migrate() error {
	log.Print("[productRepository]...Migrate")
	return p.DB.AutoMigrate(&models.Product{})
}

func (p productRepository) GetProduct(id int64, cons ...interface{}) (*models.Product, error) {
	var product models.Product
	if result := p.DB.Where("id = ?", id).First(&product); result.Error != nil {
		return &models.Product{}, result.Error
	}
	return &product, nil
}

func (p productRepository) GetProducts(cons ...interface{}) ([]*models.Product, error) {
	type query struct {
		limit int
	}
	var limit int = 8
	if cons2, ok := cons[0].(*query); ok {
		limit = cons2.limit
	}

	var products []*models.Product
	if result := p.DB.Limit(limit).Find(&products); result.Error != nil {
		return products, result.Error
	}

	return products, nil
}

func (p productRepository) CreateProduct(product *models.Product) (*models.Product, error) {
	if result := p.DB.Create(&product); result.Error != nil {
		return &models.Product{}, result.Error
	}
	return product, nil
}

func (p productRepository) UpdateProduct(product *models.Product) (*models.Product, error) {
	if result := p.DB.Save(&product); result.Error != nil {
		return &models.Product{}, result.Error
	}

	return product, nil
}

func (p productRepository) DeleteProduct(id int64) error {
	result := p.DB.Delete(&models.Product{}, id)
	return result.Error
}
