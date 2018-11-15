package router

import (
	"net/http"
	"time"
	"encoding/json"
)

type HealthCheckType struct {
	Service  string
	CurrDate string
	Version  string
}


func (r TypeRouter) HandlerHealth(w http.ResponseWriter, req *http.Request) {
	data := HealthCheckType{
		Service:  "postgres-service",
		CurrDate: time.Now().Format("2006-01-02 15:04"),
		Version:  "1.0",
	}

	payload, err := json.Marshal(data)
	if err != nil {
		r.App.Logger.Log.Error("HandlerHealth Marshal")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
