package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hoanggggg5/shopproduct/controllers/admin"
	"github.com/hoanggggg5/shopproduct/controllers/public"
	"github.com/hoanggggg5/shopproduct/controllers/resource"
	"github.com/hoanggggg5/shopproduct/params"
	"github.com/hoanggggg5/shopproduct/repositories"
	"github.com/hoanggggg5/shopproduct/routes/middlewares"
	"github.com/hoanggggg5/shopproduct/services"
	"github.com/hoanggggg5/shopproduct/usecase"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println(err)
			code := 500

			if e, ok := err.(*params.Error); ok {
				code = e.Code
			} else if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			} else {
				err = params.ServerInternalError
			}

			return c.Status(code).SendString(err.Error())
		},
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	kafkaClient := services.NewKafkaClient()

	userRepository := repositories.NewUserRepository(db)
	productRepository := repositories.NewProductRepository(db)
	categoryRepository := repositories.NewCategoryRepository(db)
	commentRepository := repositories.NewCommentRepository(db)
	orderRepository := repositories.NewOrderRepository(db)
	cartRepository := repositories.NewCartRepository(db)

	productRepository.Migrate()
	categoryRepository.Migrate()
	commentRepository.Migrate()
	orderRepository.Migrate()

	userService := usecase.NewUserService(userRepository)
	productService := usecase.NewProductService(productRepository)
	categoryService := usecase.NewCategoryService(categoryRepository)
	commentService := usecase.NewCommentService(commentRepository)
	orderService := usecase.NewOrderService(orderRepository, kafkaClient)
	cartService := usecase.NewCartService(cartRepository)

	app_api := app.Group("/api/v2")
	{
		publicRoute := app_api.Group("/public")
		adminRoute := app_api.Group("/admin")
		resourceRoute := app_api.Group("/resource")

		adminRoute.Use(middlewares.CheckRequest(userService))
		resourceRoute.Use(middlewares.CheckRequest(userService))

		public.NewRouter(publicRoute, productService, commentService)
		admin.NewRouter(adminRoute, productService, categoryService)
		resource.NewRouter(resourceRoute, commentService, orderService, cartService)
	}

	app.Listen(":3001")
}
