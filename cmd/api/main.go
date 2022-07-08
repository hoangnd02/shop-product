package main

import (
	"github.com/hoanggggg5/shopproduct/config"
	"github.com/hoanggggg5/shopproduct/models"
	"github.com/hoanggggg5/shopproduct/routes"
)

func main() {
	db := config.InitConfig()

	db.AutoMigrate(
		models.Category{},
		models.Product{},
		models.User{},
		models.Order{},
		models.ProductOrder{},
		models.ProductCart{},
		models.Comment{},
	)

	routes.InitRouter(db)
}
