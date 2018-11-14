package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"database/sql"
)

//type Dialer interface {
//	NewRouter() (Mux, error)
//}
//
//type Mux struct {
//	Router *mux.Router
//}

//func (s Mux) New() (Mux, error) {
//	return Mux{
//		Router: mux.NewRouter(),
//	}, nil
//}

type Router struct {
	//config     config.Config
	//log        logging.Logger
	app        app
	dbSession  *sql.DB
	httpClient *http.Client
}

func NewRouter() *Router {
	return &Router{
		app: app,
		//config:     config,
		//log:        log,
		dbSession:  dbSession,
	}
}

func NewRouter() *mux.Router {
	return mux.NewRouter().StrictSlash(true)
}

func (r *Router) InitMux() *mux.Router {
	r.HandleFunc("/version", rest.HandlerVersion).Methods(http.MethodGet)
	r.HandleFunc("/health", rest.HandlerHealth).Methods(http.MethodGet)
}
