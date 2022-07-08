package usecase

import (
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/repositories"
)

type ProductService interface {
	GetProduct(int64, ...interface{}) (*models.Product, error)
	GetProducts(...interface{}) ([]*models.Product, error)
	CreateProduct(*models.Product) (*models.Product, error)
	UpdateProduct(*models.Product) (*models.Product, error)
	DeleteProduct(int64) error
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(r repositories.ProductRepository) ProductService {
	return productService{
		productRepository: r,
	}
}

func (p productService) GetProduct(id int64, cons ...interface{}) (*models.Product, error) {
	return p.productRepository.GetProduct(id)
}

func (p productService) GetProducts(cons ...interface{}) ([]*models.Product, error) {
	return p.productRepository.GetProducts(cons)
}

func (p productService) CreateProduct(product *models.Product) (*models.Product, error) {
	return p.productRepository.CreateProduct(product)
}

func (p productService) UpdateProduct(product *models.Product) (*models.Product, error) {
	return p.productRepository.UpdateProduct(product)
}

func (p productService) DeleteProduct(id int64) error {
	return p.productRepository.DeleteProduct(id)
}
