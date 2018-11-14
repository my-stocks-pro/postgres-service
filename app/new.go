package app

import (
	"github.com/my-stocks-pro/postgres-service/app/config"
	"github.com/my-stocks-pro/postgres-service/app/logger"
)

type App interface {
	InitApp(c config.Config, l logger.Logger) *TypeApp
}

type TypeApp struct {
	config config.TypeConfig
	log    logger.TypeLogger
}

func NewApp() TypeApp {
	return TypeApp{}
}

func (a TypeApp) InitApp(c config.Config, l logger.Logger) *TypeApp {
	return &TypeApp{
		config: c.LoadConfig(),
		log:    l.InitLogger(),
	}
}