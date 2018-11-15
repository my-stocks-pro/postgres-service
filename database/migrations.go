package database

import (
	"github.com/jinzhu/gorm"
	"os"
	"github.com/my-stocks-pro/postgres-service/database/models"
)

func (p *Session) MakeMigrations(connection *gorm.DB) {
	if os.Getenv("MIGRATE") == "1" {
		connection.AutoMigrate(&models.Approve{}, &models.Earnings{})
		//connection.AutoMigrate()
	}
}
