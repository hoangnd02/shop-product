package resource

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/params"
	"github.com/hoanggggg5/shopproduct/pkg/entities"
)

func (h Handler) GetCart(c *fiber.Ctx) error {
	user := c.Locals("current_user").(*models.User)

	cart, err := h.cartService.GetCart(user.ID)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(cart)
}

func (h Handler) UpdateCart(c *fiber.Ctx) error {
	var payload entities.CartProductEntity
	if err := c.BodyParser(&payload); err != nil {
		return params.ServerInternalError
	}

	user := c.Locals("current_user").(*models.User)

	cart, err := h.cartService.GetCart(user.ID)
	if err != nil {
		return params.ServerInternalError
	}

	newProductCart := &models.ProductCart{
		UserID:    user.ID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
	}

	for _, productCart := range cart {
		if productCart.ProductID == payload.ProductID {
			productCart.Quantity += payload.Quantity

			productCartUpdated, err := h.cartService.UpdateCart(productCart)
			if err != nil {
				return params.ServerInternalError
			}

			return c.Status(201).JSON(productCartUpdated)
		}
	}

	newProductCart, err = h.cartService.CreateProductCart(newProductCart)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(newProductCart)
}

func (h Handler) DeleteCart(c *fiber.Ctx) error {
	idCartProduct, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	if err := h.cartService.DeleteCart(idCartProduct); err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(201)
}

func (h Handler) ClearCart(c *fiber.Ctx) error {
	user := c.Locals("current_user").(*models.User)

	if err := h.cartService.ClearCart(user.ID); err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(201)
}
