package resource

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/params"
	"github.com/hoanggggg5/shopproduct/pkg/entities"
)

func (h Handler) CreateComment(c *fiber.Ctx) error {
	idProduct, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	payload := new(entities.CommentEntity)
	if err := c.BodyParser(payload); err != nil {
		return params.ServerInternalError
	}

	user := c.Locals("current_user").(*models.User)

	comment := &models.Comment{
		UserID:    user.ID,
		ProductID: idProduct,
		Content:   payload.Content,
	}

	newComment, err := h.commentService.CreateComment(comment)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(newComment)
}

func (h Handler) UpdateComment(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return params.ServerInternalError
	}

	var payload entities.CommentEntity
	if err := c.BodyParser(&payload); err != nil {
		return params.ServerInternalError
	}

	comment, err := h.commentService.GetComment(id)
	if err != nil {
		return params.ServerInternalError
	}

	comment.Content = payload.Content

	commentUpdated, err := h.commentService.UpdateComment(comment)
	if err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(commentUpdated)
}

func (h Handler) DeleteComment(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		c.Status(500).JSON(params.ServerInternalError)
	}

	if err := h.commentService.DeleteComment(id); err != nil {
		return params.ServerInternalError
	}

	return c.Status(201).JSON(201)
}
