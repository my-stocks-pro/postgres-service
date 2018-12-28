package handler

import (
	"net/http"
	"fmt"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
	"github.com/my-stocks-pro/postgres-service/infrastructure/postgres"
)

type Default interface {
	Handle(c *gin.Context)
}

type DefaultType struct {
	config infrastructure.Config
	logger infrastructure.Logger
	redis  postgres.Postgres
}

func NewDefault(l infrastructure.Logger) DefaultType {
	return DefaultType{
		logger: l,
	}
}

func (d DefaultType) Handle(c *gin.Context) {
	paramKey := "service"
	keyRedisDB := c.Param(paramKey)
	if keyRedisDB == "" {
		d.logger.ContextError(c, http.StatusNotAcceptable, errors.New(fmt.Sprintf("Key -> %s dont exist in gin Params", keyRedisDB)))
		return
	}
	d.logger.ContextError(c, http.StatusNotAcceptable, errors.New(fmt.Sprintf("Key: %s Not Acceptable", keyRedisDB)))
}
