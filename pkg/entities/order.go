package entities

import (
	"time"

	"github.com/hoanggggg5/shopproduct/models"
)

type OrderProductEntity struct {
	ProductID int64 `json:"productId"`
	Quantity  int64 `json:"quantity"`
}

type OrderJSON struct {
	ID      int64              `json:"id"`
	Total   int64              `json:"total"`
	Address string             `json:"address"`
	Status  models.OrderStatus `json:"status"`
	// Products []*OrderProductEntity `json:"products"`
}

type OrderEntity struct {
	ID       int64                  `json:"id"`
	Total    int64                  `json:"total"`
	Address  string                 `json:"address"`
	Status   models.OrderStatus     `json:"status"`
	Products []*models.ProductOrder `json:"products"`
}

func OrderProductToEntity(productOrder *models.ProductOrder) *OrderProductEntity {
	return &OrderProductEntity{
		ProductID: productOrder.ProductID,
		Quantity:  productOrder.Quantity,
	}
}

func OrderToEntity(order *models.Order) *OrderJSON {
	// productOrders := make([]*OrderProductEntity, 0)
	// for _, p := range order.Products {
	// 	productOrders = append(productOrders, OrderProductToEntity(p))
	// }

	return &OrderJSON{
		ID:      order.ID,
		Total:   order.Total,
		Address: order.Address,
		Status:  order.Status,
	}
}

type OrderProduct struct {
	ID        int64              `json:"id"`
	Total     int64              `json:"total"`
	Address   string             `json:"address"`
	Products  int                `json:"products"`
	Status    models.OrderStatus `json:"status"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

func OrdersToEntity(orders []*models.Order) []*OrderProduct {
	ordersJSON := make([]*OrderProduct, 0)

	for _, order := range orders {
		ordersJSON = append(ordersJSON, &OrderProduct{
			ID:        order.ID,
			Address:   order.Address,
			Total:     order.Total,
			Products:  len(order.Products),
			Status:    order.Status,
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		})
	}

	return ordersJSON
}
