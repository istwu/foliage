package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID          uint    `json:"id" gorm:"primaryKey"`
	UserID      uint    `json:"user_id" gorm:"not null;constraint:OnDelete:CASCADE"`
	PortfolioID uint    `json:"portfolio_id" gorm:"not null;constraint:OnDelete:CASCADE"`
	Type        string  `json:"type" gorm:"column:post_type"`
	ImageURL    *string `json:"image_url,omitempty"`
	TextBody    *string `json:"text_body,omitempty"`
	Position    int
}
