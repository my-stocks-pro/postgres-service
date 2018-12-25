package engine

import (
	"github.com/my-stocks-pro/postgres-service/handler"
	"github.com/gin-gonic/gin"
)

const (
	version  = "version"
	health   = "health"
	earnings = "earnings"
	approved = "approved"
	rejected = "rejected"
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
	case version:
		return handler.NewVersion(s.config)
	case health:
		return handler.NewHealth(s.config, s.postgres)
	case earnings:
		return handler.NewEarnings(s.config, s.logger, s.postgres)
	case approved:
		return handler.NewApproved(s.config, s.logger, s.postgres)
	case rejected:
		return handler.NewRejected(s.config, s.logger, s.postgres)
	default:
		return handler.NewDefault(s.logger)
	}
	return nil
}