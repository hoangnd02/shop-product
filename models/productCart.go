package models

type ProductCart struct {
	ID        int64 `gorm:"primaryKey;autoIncrement;not null"`
	UserID    int64 `gorm:"type:bigint;not null;"`
	ProductID int64 `gorm:"type:bigint;not null;"`
	Quantity  int64 `gorm:"type:bigint;not null;"`
	User      *User
	Product   *Product
}
