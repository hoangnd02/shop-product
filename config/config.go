package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func createProductID(db *gorm.DB) {
// 	log.Println("hoang")
// }

func InitConfig() *gorm.DB {
	dsn := "host=localhost user=root password=123456 dbname=product port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db

	// db.Callback().Create().Register("gorm:create", createProductID)
}
