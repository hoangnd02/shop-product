package public

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/params"
)

func (h Handler) GetComments(c *fiber.Ctx) error {
	idProduct, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(500).JSON(params.ServerInternalError)
	}

	comments, err := h.commentService.GetComments(idProduct)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(comments)
}
