package models

import "github.com/jinzhu/gorm"

type Earnings struct {
	gorm.Model
	Timestamp int64  `gorm:"type:int" json:"timestamp"`
	Date      string `gorm:"size:100" json:"date"`
	IDI       int    `gorm:"size:100" json:"idi"`
	Download  int    `gorm:"size:100" json:"download"`
	Category  string `gorm:"size:100" json:"category"`
	Country   string `gorm:"size:100" json:"country"`
	City      string `gorm:"size:100" json:"city"`
}
