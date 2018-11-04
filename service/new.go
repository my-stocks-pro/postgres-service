package service

import (
	"github.com/my-stocks-pro/postgres-service/router"
	"github.com/my-stocks-pro/postgres-service/database"
)

type Service struct {
	DB     database.Conn
	Router router.Mux
}

func New() Service {
	return Service{
		DB:
	}
}




//func InitDB(p database.DB) (database.Conn, error) {
//	conn, err := p.Client()
//	if err != nil {
//		panic(err)
//	}
//
//	return conn, nil
//}
//
//func InitRouter(r router.Dialer) (router.Mux, error) {
//	mux, err := r.NewRouter()
//	if err != nil {
//		return router.Mux{}, err
//	}
//	return mux, nil
//}
