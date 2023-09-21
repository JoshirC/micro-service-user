package models

import (
	"time"

	"gorm.io/gorm"
)

type RegisterBuy struct {
	gorm.Model

	Users    Users     `gorm:"embedded"`
	IdBuy    string    `gorm:"not null;unique" json:"idBuy"`
	TotalBuy int       `gorm:"not null" json:"totalBuy"`
	DateBuy  time.Time `gorm:"not null" json:"dateBuy"`
}
