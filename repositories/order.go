package repositories

import (
	"log"

	"github.com/hoanggggg5/shopproduct/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

type OrderRepository interface {
	Migrate() error
	Find(conds ...interface{}) ([]*models.Order, error)
	First(conds ...interface{}) (*models.Order, error)
	GetOrder(int64) (*models.Order, error)
	GetOrders(int64) ([]*models.Order, error)
	GetProductOrder(int64) (*models.ProductOrder, error)
	Create(*models.Order) (*models.Order, error)
	UpdateOrder(*models.Order) (*models.Order, error)
	UpdateProductOrder(*models.ProductOrder) (*models.ProductOrder, error)
	DeleteOrder(id int64) error
	Delete(int64, interface{}) error
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return orderRepository{
		DB: db,
	}
}

func (o orderRepository) Migrate() error {
	log.Print("[OrderRepository]...Migrate")
	return o.DB.AutoMigrate(&models.Order{})
}

func (o orderRepository) First(conds ...interface{}) (order *models.Order, err error) {
	if result := o.DB.First(&order, conds...); result.Error != nil {
		return nil, result.Error
	}

	return
}

func (o orderRepository) Find(conds ...interface{}) (orders []*models.Order, err error) {
	if result := o.DB.Find(&orders, conds...); result.Error != nil {
		return nil, result.Error
	}

	return
}

func (o orderRepository) Delete(id int64, conds interface{}) (err error) {
	if result := o.DB.Where("id = ?", id).Delete(&conds); result.Error != nil {
		return result.Error
	}

	return nil
}

func (o orderRepository) GetOrder(id int64) (*models.Order, error) {
	var order *models.Order
	if result := o.DB.Preload("Products").Where("id = ?", id).Find(&order); result.Error != nil {
		return order, result.Error
	}
	return order, nil
}

func (o orderRepository) GetOrders(idUser int64) ([]*models.Order, error) {
	var orders []*models.Order
	if result := o.DB.Joins("User").Preload("Products").Where("user_id = ?", idUser).Find(&orders); result.Error != nil {
		return orders, result.Error
	}
	return orders, nil
}

func (o orderRepository) GetProductOrder(id int64) (*models.ProductOrder, error) {
	var productOrder *models.ProductOrder
	if result := o.DB.Where("id = ?", id).First(&productOrder); result.Error != nil {
		return productOrder, result.Error
	}
	return productOrder, nil
}

func (o orderRepository) Create(order *models.Order) (*models.Order, error) {
	if result := o.DB.Create(&order); result.Error != nil {
		return order, result.Error
	}
	return order, nil
}

func (o orderRepository) UpdateOrder(order *models.Order) (*models.Order, error) {
	if result := o.DB.Save(&order); result.Error != nil {
		return order, result.Error
	}
	return order, nil
}

func (o orderRepository) UpdateProductOrder(productOrder *models.ProductOrder) (*models.ProductOrder, error) {
	if result := o.DB.Save(&productOrder); result.Error != nil {
		return productOrder, result.Error
	}
	return productOrder, nil
}

func (o orderRepository) DeleteOrder(id int64) error {
	result := o.DB.Delete(&models.Order{}, id)
	return result.Error
}
