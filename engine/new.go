package engine

import (
	"github.com/my-stocks-pro/postgres-service/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/handler"
	"os"
)

type Service struct {
	config   infrastructure.Config
	logger   infrastructure.Logger
	postgres infrastructure.Postgres
	handler  map[string]handler.Handler
	QuitRPC  chan bool
	QuitTick chan bool
	QuitOS   chan os.Signal
	Engine   *gin.Engine
}

func New(c infrastructure.Config, l infrastructure.Logger, p infrastructure.Postgres) Service {
	return Service{
		config:   c,
		logger:   l,
		postgres: p,
		handler:  map[string]handler.Handler{},
		Engine:   gin.New(),
		QuitOS:   make(chan os.Signal),
		QuitRPC:  make(chan bool),
		QuitTick: make(chan bool),
	}
}
