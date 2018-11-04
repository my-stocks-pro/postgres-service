package router

import (
	"github.com/gorilla/mux"
)

type Dialer interface {
	NewRouter() (Mux, error)
}

type Mux struct {
	Router *mux.Router
}

func (s Mux) NewRouter() (Mux, error) {
	return Mux{Router: mux.NewRouter()}, nil
}
