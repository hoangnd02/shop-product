package usecase

import (
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/repositories"
)

type CartService interface {
	GetCart(int64) ([]*models.ProductCart, error)
	CreateProductCart(*models.ProductCart) (*models.ProductCart, error)
	UpdateCart(*models.ProductCart) (*models.ProductCart, error)
	DeleteCart(int64) error
	ClearCart(int64) error
}

type cartService struct {
	cartRepository repositories.CartRepository
}

func NewCartService(r repositories.CartRepository) CartService {
	return cartService{
		cartRepository: r,
	}
}

func (s cartService) GetCart(userID int64) ([]*models.ProductCart, error) {
	return s.cartRepository.GetCart(userID)
}

func (s cartService) CreateProductCart(productCart *models.ProductCart) (*models.ProductCart, error) {
	return s.cartRepository.Create(productCart)
}

func (s cartService) UpdateCart(productCart *models.ProductCart) (*models.ProductCart, error) {
	return s.cartRepository.UpdateCart(productCart)
}

func (s cartService) DeleteCart(id int64) error {
	return s.cartRepository.DeleteCart(id)
}

func (s cartService) ClearCart(userId int64) error {
	return s.cartRepository.ClearCart(userId)
}
