package models

import (
	"time"
)

type Product struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;not null"`
	Name        string    `gorm:"type:character varying;uniqueIndex:index_products_on_productname"`
	CategoryId  int64     `gorm:"type:bigint;not null;"`
	Price       int64     `gorm:"type:bigint;not null;"`
	Discount    int64     `gorm:"type:bigint;not null;"`
	Image       string    `gorm:"type:character varying;not null;"`
	Description string    `gorm:"type:character varying;"`
	CreatedAt   time.Time `gorm:"type:timestamp"`
	UpdatedAt   time.Time `gorm:"type:timestamp"`
	Category    *Category
	Comments    []*Comment
}
