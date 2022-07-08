package public

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/params"
	"github.com/hoanggggg5/shopproduct/pkg/entities"
)

func (h Handler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.FailedToParseBody
	}

	product, err := h.productService.GetProduct(id)
	if err != nil {
		return params.FailedConnectDataInDatabase
	}

	return c.Status(201).JSON(entities.ProductToEntity(product))
}

func (h Handler) GetProducts(c *fiber.Ctx) error {
	type Query struct {
		limit int
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		return params.FailedToParseBody
	}
	query := &Query{limit: limit}

	products, err := h.productService.GetProducts(query)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(entities.ProductsToEntity(products))
}
