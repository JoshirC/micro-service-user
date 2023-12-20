package models

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	Rut      string `gorm:"unique;not null" json:"rut"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"unique;not null" json:"email"`
	City     string `json:"city"`
}
