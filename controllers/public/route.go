package public

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/usecase"
)

type Handler struct {
	productService usecase.ProductService
	commentService usecase.CommentService
}

func NewRouter(
	router fiber.Router,
	productService usecase.ProductService,
	commentService usecase.CommentService,
) {

	handler := Handler{
		productService: productService,
		commentService: commentService,
	}

	router.Get("/product/:id", handler.GetProduct)
	router.Get("/products", handler.GetProducts)
	router.Get("/comments/:id", handler.GetComments)

}
