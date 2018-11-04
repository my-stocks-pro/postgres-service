package service

import (
	"github.com/my-stocks-pro/postgres-service/router"
	"github.com/gorilla/mux"
)

func (s Service) NewRouter() (router.Mux, error) {
	return router.Mux{Router: mux.NewRouter()}, nil
}
