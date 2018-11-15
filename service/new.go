package service

import (
	"github.com/my-stocks-pro/postgres-service/router"
	"github.com/my-stocks-pro/postgres-service/app"
	"net/http"
	"time"
)

type TypeService struct {
	Server *http.Server
	App    app.TypeApp
}

func NewService(router router.TypeRouter) TypeService {
	return TypeService{
		Server: &http.Server{
			Handler:      router.Router,
			Addr:         "127.0.0.1:8000",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}
