package handler

import (

	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
)

type Health interface {
	Handle(c *gin.Context)
}

type HealthType struct {
	config infrastructure.Config
	redis  infrastructure.Postgres
}

func NewHealth(c infrastructure.Config, r infrastructure.Postgres) HealthType {
	return HealthType{
		config: c,
		redis:  r,
	}
}

func (h HealthType) Handle(c *gin.Context) {

	//switch c.Request.Method {
	//case http.MethodGet:
	//	redisStatus := true
	//	if err := h.redis.Ping(); err != nil {
	//		redisStatus = false
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"service": true,
	//		"redisDB": redisStatus,
	//	})
	//default:
	//	c.JSON(http.StatusMethodNotAllowed, "Method Not Allowed")
	//}

}