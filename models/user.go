package models

import "time"

type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement;not null"`
	UID       string         `gorm:"type:character varying;uniqueIndex:index_users_on_uid"`
	Username  string         `gorm:"type:character varying;uniqueIndex:index_users_on_username"`
	Email     string         `gorm:"type:character varying;uniqueIndex:index_users_on_email"`
	Role      string         `gorm:"type:character varying;not null;default:member"`
	State     string         `gorm:"type:character varying;not null;default:pending"`
	Cart      []*ProductCart `gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time      `gorm:"type:timestamp"`
	UpdatedAt time.Time      `gorm:"type:timestamp"`
	Comments  []*Comment     `gorm:"constraint:OnDelete:CASCADE"`
	Likes     []*Like        `gorm:"constraint:OnDelete:CASCADE"`
	Orders    []*Order       `gorm:"constraint:OnDelete:CASCADE"`
}
