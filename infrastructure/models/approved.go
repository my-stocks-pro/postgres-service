package models

import "github.com/jinzhu/gorm"

type Approve struct {
	gorm.Model
	Timestamp   int64  `gorm:"type:int" json:"timestamp"`
	IDI         string `gorm:"size:30" json:"id"`
	AddedDate   string `gorm:"size:30" json:"added_date"`
	Link        string `gorm:"size:1024" json:"link"`
	Description string `gorm:"size:10240" json:"description"`
}

