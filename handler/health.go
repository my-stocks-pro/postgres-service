package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
	"github.com/my-stocks-pro/postgres-service/infrastructure/postgres"
)

type Health interface {
	Handle(c *gin.Context)
}

type HealthType struct {
	config   infrastructure.Config
	postgres postgres.Postgres
}

func NewHealth(c infrastructure.Config, p postgres.Postgres) HealthType {
	return HealthType{
		config:   c,
		postgres: p,
	}
}

func (h HealthType) Handle(c *gin.Context) {

	switch c.Request.Method {
	case http.MethodGet:
		redisStatus := true
		if err := h.postgres.Ping(); err != nil {
			redisStatus = false
		}

		c.JSON(http.StatusOK, gin.H{
			"service": true,
			"redisDB": redisStatus,
		})
	default:
		c.JSON(http.StatusMethodNotAllowed, "Method Not Allowed")
	}

}
