package service

import (
	"github.com/my-stocks-pro/postgres-service/router"
)

type Service struct {
	Conn
	Router router.Mux
}

func New() Service {
	return Service{}
}

func InitDB(p Persist) (Conn, error) {
	conf, err := p.Init()
	if err != nil {
		panic(err)
	}

	conn, err := p.Client(conf)
	if err != nil {
		panic(err)
	}

	return conn, nil
}

func InitRouter(r router.Dialer) (router.Mux, error) {
	mux, err := r.NewRouter()
	if err != nil {
		return router.Mux{}, err
	}
	return mux, nil
}