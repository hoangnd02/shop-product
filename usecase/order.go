package usecase

import (
	"encoding/json"

	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/params"
	"github.com/hoanggggg5/shopproduct/pkg/entities"
	"github.com/hoanggggg5/shopproduct/repositories"
	"github.com/hoanggggg5/shopproduct/services"
	"github.com/twmb/franz-go/pkg/kgo"
)

type OrderService interface {
	GetOrders(id int64) ([]*models.Order, error)
	GetOrder(id int64) (*models.Order, error)
	GetProductOrder(int64) (*models.ProductOrder, error)
	Create(*models.Order) (*models.Order, error)
	UpdateInfoOrder(*models.Order) (*models.Order, error)
	UpdateProductOrder(*models.ProductOrder) (*models.ProductOrder, error)
	DeleteProductOrder(int64) error
	DeleteOrder(int64) error
	SendMail(*entities.MailerPayload) error
}

type orderService struct {
	orderRepository repositories.OrderRepository
	kafkaClient     *services.KafkaClient
}

func NewOrderService(r repositories.OrderRepository, kafkaClient *services.KafkaClient) OrderService {
	return orderService{
		orderRepository: r,
		kafkaClient:     kafkaClient,
	}
}

func (s orderService) GetOrders(idUser int64) ([]*models.Order, error) {
	return s.orderRepository.GetOrders(idUser)
}

func (s orderService) GetOrder(id int64) (*models.Order, error) {
	return s.orderRepository.GetOrder(id)
}

func (s orderService) GetProductOrder(id int64) (*models.ProductOrder, error) {
	return s.orderRepository.GetProductOrder(id)
}

func (s orderService) Create(order *models.Order) (*models.Order, error) {
	return s.orderRepository.Create(order)
}

func (s orderService) UpdateInfoOrder(order *models.Order) (*models.Order, error) {
	return s.orderRepository.UpdateOrder(order)
}

func (s orderService) UpdateProductOrder(productOrder *models.ProductOrder) (*models.ProductOrder, error) {
	return s.orderRepository.UpdateProductOrder(productOrder)
}

func (s orderService) DeleteProductOrder(id int64) error {
	var productOrder *models.ProductOrder
	return s.orderRepository.Delete(id, productOrder)
}

func (s orderService) DeleteOrder(id int64) error {
	var order *models.Order
	return s.orderRepository.Delete(id, order)
}

func (s orderService) SendMail(mailerPayload *entities.MailerPayload) error {
	message := mailerPayload

	messageJson, err := json.Marshal(message)
	if err != nil {
		return params.ServerInternalError
	}

	record := &kgo.Record{
		Topic: "mailer",
		Value: []byte(messageJson),
	}

	s.kafkaClient.Produce(record)
	return nil
}
