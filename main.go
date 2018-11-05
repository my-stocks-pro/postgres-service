package main

import (
	"fmt"
	"github.com/my-stocks-pro/postgres-service/service"
	"github.com/my-stocks-pro/postgres-service/config"
)

func main() {
	fmt.Println("POSTGRES")

	config, err := config.New().Load()
	if err != nil {
		fmt.Println(err)
	}

	srv := service.New()



}
