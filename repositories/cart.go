package repositories

import (
	"github.com/hoanggggg5/shopproduct/models"
	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

type CartRepository interface {
	GetCart(int64) ([]*models.ProductCart, error)
	Create(*models.ProductCart) (*models.ProductCart, error)
	UpdateCart(*models.ProductCart) (*models.ProductCart, error)
	DeleteCart(id int64) error
	ClearCart(id int64) error
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return cartRepository{
		DB: db,
	}
}

func (c cartRepository) GetCart(userID int64) ([]*models.ProductCart, error) {
	var carts []*models.ProductCart
	if result := c.DB.Where("user_id = ?", userID).Find(&carts); result.Error != nil {
		return carts, result.Error
	}
	return carts, nil
}

func (c cartRepository) Create(cart *models.ProductCart) (*models.ProductCart, error) {
	if result := c.DB.Create(&cart); result.Error != nil {
		return cart, result.Error
	}
	return cart, nil
}

func (c cartRepository) UpdateCart(cart *models.ProductCart) (*models.ProductCart, error) {
	if result := c.DB.Save(&cart); result.Error != nil {
		return cart, result.Error
	}
	return cart, nil
}

func (c cartRepository) DeleteCart(id int64) error {
	var cartProduct *models.ProductCart
	if result := c.DB.Delete(&cartProduct, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (c cartRepository) ClearCart(id int64) error {
	result := c.DB.Where("user_id = ?", id).Delete(&models.ProductCart{}, id)
	return result.Error
}
