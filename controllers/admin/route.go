package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/usecase"
)

type Handler struct {
	productService  usecase.ProductService
	categoryService usecase.CategoryService
}

func NewRouter(
	router fiber.Router,
	productService usecase.ProductService,
	categoryService usecase.CategoryService,
) {
	handler := Handler{
		productService:  productService,
		categoryService: categoryService,
	}

	router.Post("/product", handler.CreateProduct)
	router.Put("/product/:id", handler.UpdateProduct)
	router.Delete("/product/:id", handler.DeleteProduct)

	router.Get("/categories", handler.GetCategoies)
}
