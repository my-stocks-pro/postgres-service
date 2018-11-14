package app

import "github.com/my-stocks-pro/postgres-service/config"

type App struct {
	config     config
	log      log
}

func NewApp(c config, l logger) *App {
	return &App{
		config: c,
		log:    l,
	}
}
