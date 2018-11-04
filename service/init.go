package service

import (
	"os"
)

type Config struct {
	HOST string
	PORT string
	NAME string
	USER string
	PASS string
}

func (s Service) Init() (Config, error) {
	return Config{
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("NAME"),
		os.Getenv("USER"),
		os.Getenv("PASS"),
	}, nil
}
