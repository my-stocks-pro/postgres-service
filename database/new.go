package database

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Persist interface {
	NewClient() Session
}

type Session struct {
	Postgres *gorm.DB
}

func NewSession() Session {
	return Session{}
}


func (с Session) NewClient() Session {
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		c.HOST, c.PORT, c.NAME, c.USER, c.PASS)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		return Conn{}, err
	}

	err = connection.DB().Ping()
	if err != nil {
		return Conn{}, err
	}

	//p.MakeMigrations(connection)

	return Conn{Postgres: connection}, nil

}
