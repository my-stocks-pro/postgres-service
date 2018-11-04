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


