package models

import "gorm.io/gorm"

type Portfolio struct {
	gorm.Model
	ID     uint   `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	UserID uint   `json:"user_id" gorm:"not null;constraint:OnDelete:CASCADE"`
	Posts  []Post `gorm:"foreignKey:PortfolioID"`
}
