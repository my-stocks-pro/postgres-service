package database

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Persist interface {
	NewClient() (Conn, error)
}

type DB struct {

}

type Conn struct {
	Postgres *gorm.DB
}

func (d DB) NewClient() (Conn, error) {
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
