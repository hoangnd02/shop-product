package resource

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/params"
	"github.com/hoanggggg5/shopproduct/pkg/entities"
)

func (h Handler) Create(c *fiber.Ctx) error {
	payload := new(entities.OrderEntity)
	if err := c.BodyParser(payload); err != nil {
		return params.ServerInternalError
	}

	user := c.Locals("current_user").(*models.User)

	order := &models.Order{
		UserID:  user.ID,
		Total:   payload.Total,
		Address: payload.Address,
	}

	newOrder, err := h.orderService.Create(order)
	if err != nil {
		return params.ServerInternalError
	}

	newOrder.Products = payload.Products

	updatedOrder, err := h.orderService.UpdateInfoOrder(newOrder)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(entities.OrderToEntity(updatedOrder))
}

func (h Handler) GetOrders(c *fiber.Ctx) error {
	idUser := c.Locals("current_user").(*models.User).ID
	orders, err := h.orderService.GetOrders(idUser)

	if err != nil {
		return c.Status(422).JSON(params.FailedConnectDataInDatabase)
	}

	return c.Status(201).JSON(entities.OrdersToEntity(orders))
}

func (h Handler) GetOrder(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	order, err := h.orderService.GetOrder(id)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(entities.OrderToEntity(order))
}

func (h Handler) UpdateInfoOrder(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	var payload entities.OrderEntity
	if err := c.BodyParser(&payload); err != nil {
		return params.ServerInternalError
	}

	order, err := h.orderService.GetOrder(id)
	if err != nil {
		return params.ServerInternalError
	}

	order.Address = payload.Address
	order.Total = payload.Total
	order.Status = payload.Status
	order.Products = payload.Products

	orderUpdated, err := h.orderService.UpdateInfoOrder(order)
	if err != nil {
		return err
	}

	if orderUpdated.Status == models.OrderStatusDelivering {
		message := &entities.MailerPayload{
			Key: "status.order",
			To:  order.Address,
			Record: map[string]interface{}{
				"Status": "delivering",
			},
		}

		h.orderService.SendMail(message)
	}

	return c.Status(201).JSON(entities.OrderToEntity(orderUpdated))
}

func (h Handler) UpdateProductOrder(c *fiber.Ctx) error {
	idProductOrder, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	var payload entities.OrderProductEntity
	if err := c.BodyParser(&payload); err != nil {
		return params.ServerInternalError
	}

	productOrder, err := h.orderService.GetProductOrder(idProductOrder)
	if err != nil {
		return params.ServerInternalError
	}

	productOrder.Quantity = payload.Quantity

	productOrderUpdated, err := h.orderService.UpdateProductOrder(productOrder)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(productOrderUpdated)
}

func (h Handler) DeleteProductOrder(c *fiber.Ctx) error {
	idProductOrder, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	if err := h.orderService.DeleteProductOrder(idProductOrder); err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(201)
}

func (h Handler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	if err := h.orderService.DeleteOrder(id); err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(201)
}
