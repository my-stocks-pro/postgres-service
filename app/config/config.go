package config

import "os"

type Config interface {
	LoadConfig() (TypeConfig, error)
}

type TypeConfig struct {
	HOST string
	PORT string
	NAME string
	USER string
	PASS string
}

func NewConfig() TypeConfig {
	return TypeConfig{}
}

func (c TypeConfig) LoadConfig() (TypeConfig, error) {
	return TypeConfig{
		HOST: os.Getenv("HOST"),
		PORT: os.Getenv("PORT"),
		NAME: os.Getenv("NAME"),
		USER: os.Getenv("USER"),
		PASS: os.Getenv("PASS"),
	}, nil
}