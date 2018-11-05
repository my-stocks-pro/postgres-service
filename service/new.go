package service

import (
	"github.com/my-stocks-pro/postgres-service/router"
	"github.com/my-stocks-pro/postgres-service/database"
)

type Srvice interface {

}


type TypeService struct {
	DB     database.Conn
	Router router.Mux
}

func New() TypeService {
	return TypeService{}
}
