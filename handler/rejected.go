package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
)

type Rejected interface {
	Handle(c *gin.Context)
}

type TypeRejected struct {
	config   infrastructure.Config
	logger   infrastructure.Logger
	postgres infrastructure.Postgres
}

func NewRejected(c infrastructure.Config, l infrastructure.Logger, p infrastructure.Postgres) TypeRejected {
	return TypeRejected{
		config:   c,
		logger:   l,
		postgres: p,
	}
}

func (s TypeRejected) Handle(c *gin.Context) {

}
