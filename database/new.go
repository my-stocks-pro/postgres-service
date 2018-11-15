package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"github.com/my-stocks-pro/postgres-service/app"
)

type Persist interface {
	NewClient() (Session, error)
}

type Session struct {
	Client *gorm.DB
	App app.TypeApp
}

func NewSession(a app.TypeApp) Session {
	return Session{
		App: a,
	}
}

func (s Session) NewClient() (Session, error) {
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		s.App.Config.HOST, s.App.Config.PORT, s.App.Config.NAME, s.App.Config.USER, s.App.Config.PASS)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		return Session{}, err
	}

	err = connection.DB().Ping()
	if err != nil {
		return Session{}, err
	}

	s.Client = connection

	s.MakeMigrations(connection)

	return s, nil
}
