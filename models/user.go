package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	// Email    string `gorm:"unique;not null"`
	// PasswordHash string `gorm:"not null"`
	Role       string      `json:"role" gorm:"default:user"`
	Portfolios []Portfolio `gorm:"foreignKey:UserID"`
	Posts      []Post      `gorm:"foreignKey:UserID"`
}
