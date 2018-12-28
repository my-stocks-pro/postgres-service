package engine

import (
	"github.com/my-stocks-pro/postgres-service/handler"
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
)

func (s *Service) InitMux() {
	s.Engine.GET("/:service", s.GetHandler)
	s.Engine.POST("/:service", s.GetHandler)
	s.Engine.DELETE("/:service", s.GetHandler)
}

func (s *Service) GetHandler(c *gin.Context) {
	serviceType := c.Param("service")
	_, ok := s.handler[serviceType]
	if !ok {
		s.handler[serviceType] = s.HandlerConstruct(serviceType)
	}
	s.handler[serviceType].Handle(c)
}

func (s *Service) HandlerConstruct(serviceType string) handler.Handler {
	switch serviceType {
	case infrastructure.Version:
		return handler.NewVersion(s.config)
	case infrastructure.Health:
		return handler.NewHealth(s.config, s.postgres)
	case infrastructure.Earnings, infrastructure.Approved, infrastructure.Rejected:
		return handler.NewCommon(s.config, s.logger, s.postgres)
	default:
		return handler.NewDefault(s.logger)
	}
	return nil
}
