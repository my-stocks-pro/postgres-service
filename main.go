package main

import (
	"fmt"
	"github.com/my-stocks-pro/postgres-service/service"
)

func main() {
	fmt.Println("POSTGRES")

	srv := service.New()

	service.InitRouter(srv)

	db, err := service.InitDB(srv)
	if err != nil {
		panic(err)
	}



	fmt.Println(db.Postgres)

}
