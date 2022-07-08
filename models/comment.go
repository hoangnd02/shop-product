package models

import "time"

type Comment struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null"`
	UserID    int64     `gorm:"type:bigint;not null;"`
	ProductID int64     `gorm:"type:character varying;not null;"`
	Content   string    `gorm:"type:character varying;not null"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	Likes     []*Like   `json:"-" gorm:"constraint:OnDelete:CASCADE"`
	User      *User
	Product   *Product
}
