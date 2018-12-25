package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
	"net/http"
	"errors"
	"io/ioutil"
	"encoding/json"
	"io"
	"github.com/my-stocks-pro/postgres-service/infrastructure/models"
)

type Earnings interface {
	Handle(c *gin.Context)
}

type TypeEarnings struct {
	config   infrastructure.Config
	logger   infrastructure.Logger
	postgres infrastructure.Postgres
}

func NewEarnings(c infrastructure.Config, l infrastructure.Logger, p infrastructure.Postgres) TypeEarnings {
	return TypeEarnings{
		config:   c,
		logger:   l,
		postgres: p,
	}
}

func (e TypeEarnings) Handle(c *gin.Context) {
	db := c.Param("service")

	body, err := readBody(c.Request.Body)
	if err != nil {
		e.logger.ContextError(c, http.StatusInternalServerError, err)
		return
	}

	switch c.Request.Method {
	case http.MethodGet:
		val, err := e.postgres
		if err != nil {
			e.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

		_, err = c.Writer.Write(val)
		if err != nil {
			e.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

	case http.MethodPost:

		body.Update(e.postgres)


		if err := e.postgres.Update(); err != nil {
			e.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

	case http.MethodDelete:
		err := e.postgres.Delete()
		if err != nil {
			e.logger.ContextError(c, http.StatusInternalServerError, err)
			return
		}

	default:
		e.logger.ContextError(c, http.StatusMethodNotAllowed, errors.New("Method Not Allowed"))
		return
	}

	e.logger.ContextSuccess(c, http.StatusOK)
}

func readBody(b io.Reader) (*models.Earnings, error) {
	res := new(models.Earnings)
	body, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}
