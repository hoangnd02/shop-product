package models

import "time"

type Like struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;not null"`
	UserID    int64     `gorm:"type:bigint;not null;"`
	CommentID int64     `gorm:"type:bigint;not null;"`
	ProductID int64     `gorm:"type:bigint;not null;"`
	CreatedAt time.Time `gorm:"type:timestamp"`
	UpdatedAt time.Time `gorm:"type:timestamp"`
	User      *User
	Product   *Product
}
