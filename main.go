package main

import (
	"fmt"
	"github.com/my-stocks-pro/postgres-service/app"
	"github.com/my-stocks-pro/postgres-service/router"
	"github.com/my-stocks-pro/postgres-service/service"
	"github.com/my-stocks-pro/postgres-service/database"
	"github.com/my-stocks-pro/postgres-service/app/config"
	"github.com/my-stocks-pro/postgres-service/app/logger"
)

func main() {
	fmt.Println("POSTGRES")

	conf := config.NewConfig()
	log := logger.NewLogger()

	a := app.NewApp().InitApp(conf, log)

	db, err := database.NewSession(a).NewClient()
	if err != nil {
		a.Logger.Log.Error(err.Error())
	}

	r := router.NewRouter(a, db).InitMux()

	s := service.NewService(r)

	fmt.Println("SUCCESS")

	if err := s.Server.ListenAndServe(); err != nil {
		s.App.Logger.Log.Error(err.Error())
	}
}
