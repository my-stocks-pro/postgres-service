package main

import (
	"fmt"
	"github.com/my-stocks-pro/postgres-service/router"
)

func main() {
	fmt.Println("POSTGRES")
	//
	//config, err := config.New().Load()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(config)

	//srv := service.New()

	router := router.New()



	fmt.Println("SUCCESS")

}
