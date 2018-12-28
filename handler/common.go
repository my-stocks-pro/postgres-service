package handler

import (
	"net/http"

	"github.com/my-stocks-pro/postgres-service/infrastructure/postgres"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Common interface {
	Handle(c *gin.Context)
}

type TypeCommon struct {
	config   infrastructure.Config
	logger   infrastructure.Logger
	postgres postgres.Postgres
	persist  map[string]postgres.Persist
}

func NewCommon(c infrastructure.Config, l infrastructure.Logger, p postgres.Postgres) TypeCommon {
	return TypeCommon{
		config:   c,
		logger:   l,
		postgres: p,
		persist:  map[string]postgres.Persist{},
	}
}

func (h TypeCommon) Handle(c *gin.Context) {

	serviceName := c.Param("service")

	persist := h.Decorator(serviceName)
	if persist == nil {
		h.logger.ContextError(c, http.StatusInternalServerError, errors.Errorf("%s failed persist decorator", serviceName))
		return
	}

	switch c.Request.Method {
	case http.MethodGet:
		val, err := persist.Select(c.Request.Body)
		if err != nil {
			h.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

		_, err = c.Writer.Write(val)
		if err != nil {
			h.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}
	case http.MethodPost:
		if err := persist.Update(c.Request.Body); err != nil {
			h.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}
	case http.MethodDelete:
		if err := persist.Delete(c.Request.Body); err != nil {
			h.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}
	default:
		h.logger.ContextError(c, http.StatusMethodNotAllowed, errors.Errorf("%s Method Not Allowed", serviceName))
		return
	}

	h.logger.ContextSuccess(c, http.StatusOK)
}

func (h TypeCommon) Decorator(serviceName string) postgres.Persist {
	_, ok := h.persist[serviceName]
	if !ok {
		h.persist[serviceName] = h.PersistConstruct(serviceName)
	}
	return h.persist[serviceName]
}

func (h *TypeCommon) PersistConstruct(serviceType string) postgres.Persist {
	switch serviceType {
	case infrastructure.Earnings:
		return postgres.NewEarnings(h.postgres)
	case infrastructure.Approved:
		return postgres.NewApproved(h.postgres)
	case infrastructure.Rejected:
		return postgres.NewRejected(h.postgres)
	default:
		return nil
	}
	return nil
}
