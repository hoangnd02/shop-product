package resource

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/shopproduct/usecase"
)

type Handler struct {
	commentService usecase.CommentService
	orderService   usecase.OrderService
	cartService    usecase.CartService
}

func NewRouter(
	router fiber.Router,
	commentService usecase.CommentService,
	orderService usecase.OrderService,
	cartService usecase.CartService,
) {

	handler := Handler{
		commentService: commentService,
		orderService:   orderService,
		cartService:    cartService,
	}

	router.Post("/comment/:id", handler.CreateComment)

	router.Post("/order", handler.Create)
	router.Get("/orders/", handler.GetOrders)
	router.Get("/order/:id", handler.GetOrder)
	router.Put("/order/:id", handler.UpdateInfoOrder)
	router.Put("/orderproduct/:id", handler.UpdateProductOrder)
	router.Delete("/orderproduct/:id", handler.DeleteProductOrder)
	router.Delete("/order/:id", handler.DeleteOrder)

	router.Get("/cart", handler.GetCart)
	router.Post("/cart", handler.UpdateCart)
	router.Delete("/cart", handler.ClearCart)
	router.Delete("/cart/:id", handler.DeleteCart)

}
