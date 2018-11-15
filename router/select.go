package router

import (
	"net/http"
	"io/ioutil"
	"github.com/my-stocks-pro/postgres-service/database/models"
	"encoding/json"
)

type Response struct {
	Error   error
	Type    string
	Data    []byte
	Message string
}

func (r TypeRouter) Select(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")

	reqType := req.Header.Get("type")
	switch reqType {
	case "approved":
		w.Write(r.approvedSelect(body))
	case "earnings":
		w.Write(r.earningsSelect(body))
	default:
		w.Write(newResponse(nil, reqType, []byte{}, "Not allowed type"))
	}
}

func (r TypeRouter) approvedSelect(body []byte) []byte {
	approved := models.Approve{}
	r.DB.Client.Find(&approved)
	blob, err := json.Marshal(approved)
	if err != nil {
		return newResponse(err, "approved", []byte{}, "error")
	}
	return newResponse(nil, "approved", blob, "all")
}

func (r TypeRouter) earningsSelect(body []byte) []byte {
	earning := models.Earnings{}
	r.DB.Client.Find(&earning)
	blob, err := json.Marshal(earning)
	if err != nil {
		return newResponse(err, "earnings", []byte{}, "error")
	}
	return newResponse(nil, "earnings", blob, "all")
}

func newResponse(err error, reqType string, data []byte, message string) []byte {
	resp := Response{
		Error:   err,
		Type:    reqType,
		Data:    data,
		Message: message,
	}
	blob, err := json.Marshal(resp)
	if err != nil {

	}
	return blob
}
