package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	Rut      string `gorm:"not null;unique" json:"rut"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"not null" json:"email"`
	City     string `gorm:"not null" json:"city"`
}
