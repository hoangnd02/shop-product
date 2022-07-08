package repositories

import (
	"log"

	"github.com/hoanggggg5/shopproduct/models"
	"gorm.io/gorm"
)

type categoryRepository struct {
	DB *gorm.DB
}

type CategoryRepository interface {
	Migrate() error
	Find([]*models.Category) ([]*models.Category, error)
	First(*models.Category) (*models.Category, error)
	Create(models.Category) (models.Category, error)
	Update(models.Category) (models.Category, error)
	Delete(models.Category) error
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return categoryRepository{
		DB: db,
	}
}

func (p categoryRepository) Migrate() error {
	log.Print("[categoryRepository]...Migrate")
	return p.DB.AutoMigrate(&models.Category{})
}

func (p categoryRepository) Find(categories []*models.Category) ([]*models.Category, error) {
	result := p.DB.Find(&categories)
	return categories, result.Error
}

func (p categoryRepository) First(category *models.Category) (*models.Category, error) {
	result := p.DB.First(&category)
	return category, result.Error
}

func (p categoryRepository) Create(category models.Category) (models.Category, error) {
	result := p.DB.Create(&category)
	return category, result.Error
}

func (p categoryRepository) Update(category models.Category) (models.Category, error) {
	result := p.DB.Save(&category)
	return category, result.Error
}

func (p categoryRepository) Delete(category models.Category) error {
	result := p.DB.Delete(&models.Category{}, category.ID)
	return result.Error
}
