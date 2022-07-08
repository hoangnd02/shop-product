package usecase

import (
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/repositories"
)

type CategoryService interface {
	GetCategories() ([]*models.Category, error)
	GetCategory() (*models.Category, error)
	CreateCategory() (models.Category, error)
}

type categoryService struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryService(r repositories.CategoryRepository) CategoryService {
	return categoryService{
		categoryRepository: r,
	}
}

func (s categoryService) GetCategories() ([]*models.Category, error) {
	var categories []*models.Category
	return s.categoryRepository.Find(categories)
}

func (s categoryService) GetCategory() (*models.Category, error) {
	var category *models.Category
	return s.categoryRepository.First(category)
}

func (s categoryService) CreateCategory() (models.Category, error) {
	var category models.Category

	return s.categoryRepository.Create(category)
}
