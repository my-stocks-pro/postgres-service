package main

import (
	"fmt"
	"github.com/my-stocks-pro/postgres-service/app"
	"github.com/my-stocks-pro/postgres-service/router"
	"github.com/my-stocks-pro/postgres-service/service"
	"github.com/my-stocks-pro/postgres-service/database"
)

func main() {
	fmt.Println("POSTGRES")

	a := app.NewApp()

	r := router.NewRouter()

	db := database.NewSession()

	srv := service.NewService(a, r, db)

	fmt.Println("SUCCESS")

}
