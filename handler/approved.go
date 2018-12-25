package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
)

type Approved interface {
	Handle(c *gin.Context)
}

type TypeApproved struct {
	config   infrastructure.Config
	logger   infrastructure.Logger
	postgres infrastructure.Postgres
}

func NewApproved(c infrastructure.Config, l infrastructure.Logger, p infrastructure.Postgres) TypeApproved {
	return TypeApproved{
		config:   c,
		logger:   l,
		postgres: p,
	}
}

func (s TypeApproved) Handle(c *gin.Context) {

}
