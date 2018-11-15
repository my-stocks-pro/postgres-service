package router

import (
	"net/http"
	"github.com/my-stocks-pro/postgres-service/database/models"
	"io/ioutil"
	"encoding/json"
)

func (r TypeRouter) Save(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	switch req.Header.Get("type") {
	case "approved":
		r.approvedSave(body)
	case "earnings":
		r.earningsSave(body)
	default:
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Not allowed type"))
	}
}

func (r TypeRouter) approvedSave(body []byte) {
	approved := models.Approve{}
	if err := json.Unmarshal(body, &approved); err != nil {
		r.App.Logger.Log.Error(err.Error())
	}
	r.DB.Client.Where("id_i=?", approved.IDI).Assign(approved).FirstOrCreate(&approved)
}

func (r TypeRouter) earningsSave(body []byte) {
	earnings := models.Earnings{}
	if err := json.Unmarshal(body, &earnings); err != nil {
		r.App.Logger.Log.Error(err.Error())
	}
	r.DB.Client.Where("id_i=?", earnings.IDI).Omit("country", "city").Assign(earnings).FirstOrCreate(&earnings)
}
