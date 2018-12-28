package infrastructure

import (
	"time"
	"os"
)

const (
	Version  = "version"
	Health   = "health"
	Earnings = "earnings"
	Approved = "approved"
	Rejected = "rejected"
	logPath   = "app_log"
	logPrefix = "postgres-service"
)

type Config struct {
	StartDate string
	SName     string
	SPort     string
	PHOST     string
	PPORT     string
	PNAME     string
	PUSER     string
	PPASS     string
}

func NewConfig() Config {
	return Config{
		StartDate: time.Now().Format("2006-01-02 15:04"),
		SName:     "postgres-service",
		SPort:     ":9006",
		PHOST:     os.Getenv("HOST"),
		PPORT:     os.Getenv("PORT"),
		PNAME:     os.Getenv("NAME"),
		PUSER:     os.Getenv("USER"),
		PPASS:     os.Getenv("PASS"),
	}
}
