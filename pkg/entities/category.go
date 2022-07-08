package entities

import (
	"github.com/hoanggggg5/shopproduct/models"
)

type CategoryEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func CategoryToEntity(category *models.Category) *CategoryEntity {
	return &CategoryEntity{
		ID:   category.ID,
		Name: category.Name,
	}
}
