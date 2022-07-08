package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/params"
)

func (h Handler) GetCategoies(c *fiber.Ctx) error {
	categories, err := h.categoryService.GetCategories()
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(categories)
}
