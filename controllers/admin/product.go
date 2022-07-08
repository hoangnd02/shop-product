package admin

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/pkg/entities"
)

func (h Handler) CreateProduct(c *fiber.Ctx) error {
	var payload entities.ProductEntity
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	product := &models.Product{
		Name:        payload.Name,
		CategoryId:  payload.CategoryID,
		Price:       payload.Price,
		Discount:    payload.Discount,
		Description: payload.Description,
		Image:       payload.Image,
	}

	newProduct, err := h.productService.CreateProduct(product)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(entities.ProductToEntity(newProduct))
}

func (h Handler) UpdateProduct(c *fiber.Ctx) error {
	idProduct, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	var payload entities.ProductEntity
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	product, err := h.productService.GetProduct(idProduct)
	if err != nil {
		return err
	}

	product.Name = payload.Name
	product.CategoryId = payload.CategoryID
	product.Price = payload.Price
	product.Discount = payload.Discount
	product.Description = payload.Description
	product.Image = payload.Image

	productUpdated, err := h.productService.UpdateProduct(product)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(entities.ProductToEntity(productUpdated))
}

func (h Handler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	return h.productService.DeleteProduct(id)
}
