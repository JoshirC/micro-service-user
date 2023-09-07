package models

import (
	"time"

	"gorm.io/gorm"
)

type RegisterBuyUser struct {
	gorm.Model

	Id       uint      `gorm:"not null; primaryKey" json:"id"`
	User     User      `gorm:"embedded"`
	IdBuy    string    `gorm:"not null" json:"idBuy"`
	TotalBuy int       `gorm:"not null" json:"totalBuy"`
	RutUser  string    `gorm:"not null" json:"rutUser"`
	DateBuy  time.Time `gorm:"not null" json:"dateBuy"`
}
