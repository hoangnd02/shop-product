package entities

import (
	"github.com/hoanggggg5/shopproduct/models"
)

type ProductEntity struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	CategoryID  int64  `json:"categoryId"`
	Price       int64  `json:"price"`
	Discount    int64  `json:"discount"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func ProductToEntity(product *models.Product) *ProductEntity {
	return &ProductEntity{
		ID:          product.ID,
		Name:        product.Name,
		CategoryID:  product.CategoryId,
		Price:       product.Price,
		Discount:    product.Discount,
		Description: product.Description,
		Image:       product.Image,
	}
}

func ProductsToEntity(products []*models.Product) []*ProductEntity {
	productEntities := make([]*ProductEntity, 0)

	for _, p := range products {
		productEntities = append(productEntities, ProductToEntity(p))
	}

	return productEntities
}
