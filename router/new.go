package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/my-stocks-pro/postgres-service/app"
	"github.com/my-stocks-pro/postgres-service/database"
)

type Router interface {
	InitMux() TypeRouter
}

type TypeRouter struct {
	App    app.TypeApp
	Router *mux.Router
	DB     database.Session
}

func NewRouter(app app.TypeApp, db database.Session) TypeRouter {
	router := mux.NewRouter().StrictSlash(true)
	return TypeRouter{
		App:    app,
		Router: router,
		DB:     db,
	}
}

func (r TypeRouter) InitMux() TypeRouter {
	r.Router.HandleFunc("/health", r.HandlerHealth).Methods(http.MethodGet)

	r.Router.HandleFunc("/postgres", r.Save).Methods(http.MethodPost)
	r.Router.HandleFunc("/postgres", r.Select).Methods(http.MethodGet)

	return TypeRouter{Router: r.Router}
}
