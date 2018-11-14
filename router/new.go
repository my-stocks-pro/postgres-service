package router

import (
	"github.com/gorilla/mux"
)

type Router interface {
	InitMux()
}

type TypeRouter struct {
	Router *mux.Router
}


func NewRouter() TypeRouter {
	return TypeRouter{
		Router: mux.NewRouter().StrictSlash(true),
	}
}

func (r TypeRouter) InitMux() {
	//r.HandleFunc("/version", rest.HandlerVersion).Methods(http.MethodGet)
	//r.HandleFunc("/health", rest.HandlerHealth).Methods(http.MethodGet)
}
