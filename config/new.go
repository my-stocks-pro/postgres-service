package config

import "os"

type Config interface {
	Load() (TypeConfig, error)
}

type TypeConfig struct {
	HOST string
	PORT string
	NAME string
	USER string
	PASS string
}

func (c TypeConfig) Load() (TypeConfig, error) {
	return TypeConfig{
		HOST: os.Getenv("HOST"),
		PORT: os.Getenv("PORT"),
		NAME: os.Getenv("NAME"),
		USER: os.Getenv("USER"),
		PASS: os.Getenv("PASS"),
	}, nil
}

func New() TypeConfig {
	return TypeConfig{}
}