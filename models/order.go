package models

import (
	"time"
)

type ProductOrder struct {
	ID        int64 `gorm:"primaryKey;autoIncrement;not null"`
	OrderID   int64 `gorm:"type:bigint;not null;"`
	ProductID int64 `gorm:"type:bigint;not null;"`
	Quantity  int64 `gorm:"type:bigint;not null;"`
	Product   *Product
}

type OrderStatus string

const (
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusDelivering OrderStatus = "delivering"
)

type Order struct {
	ID        int64           `gorm:"primaryKey;autoIncrement;not null"`
	UserID    int64           `gorm:"type:bigint;not null;"`
	Total     int64           `gorm:"type:bigint;not null;"`
	Address   string          `gorm:"type:character varying;not null;"`
	Status    OrderStatus     `gorm:"type:character varying;not null;default:'pending'"`
	CreatedAt time.Time       `gorm:"type:timestamp"`
	UpdatedAt time.Time       `gorm:"type:timestamp"`
	Products  []*ProductOrder `gorm:"constraint:OnDelete:CASCADE"`
	User      *User
}
