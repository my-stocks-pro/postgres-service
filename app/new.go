package app

import (
	"github.com/my-stocks-pro/postgres-service/app/config"
	"github.com/my-stocks-pro/postgres-service/app/logger"
)

type App interface {
	InitApp(c config.Config, l logger.Logger) *TypeApp
}

type TypeApp struct {
	Config config.TypeConfig
	Logger logger.TypeLogger
}

func NewApp() TypeApp {
	return TypeApp{}
}

func (a TypeApp) InitApp(c config.Config, l logger.Logger) TypeApp {
	return TypeApp{
		Config: c.LoadConfig(),
		Logger: l.InitLogger(),
	}
}
