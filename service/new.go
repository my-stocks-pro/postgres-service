package service

import (
	"github.com/my-stocks-pro/postgres-service/router"
	"github.com/my-stocks-pro/postgres-service/database"
	"github.com/my-stocks-pro/postgres-service/app"
)

type Srvice interface {
}

type TypeService struct {
	app    app.TypeApp
	router router.TypeRouter
	db     database.Session
}

func NewService(app app.App, router router.Router, db database.Persist) TypeService {
	return TypeService{
		app: app.InitApp()
	}
}
