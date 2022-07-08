package models

import "time"

type Category struct {
	ID        int64      `gorm:"primaryKey;autoIncrement;not null"`
	Name      string     `gorm:"type:character varying;uniqueIndex:index_products_on_productname"`
	CreatedAt time.Time  `gorm:"type:timestamp"`
	UpdatedAt time.Time  `gorm:"type:timestamp"`
	Products  []*Product `gorm:"constraint:OnDelete:CASCADE"`
}
