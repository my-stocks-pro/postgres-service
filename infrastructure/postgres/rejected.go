package postgres

import (
	"github.com/jinzhu/gorm"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type RejectedRecord struct {
	gorm.Model
	IDI int `gorm:"size:100" json:"idi"`
}

type Rejected struct {
	record   RejectedRecord
	postgres Postgres
}

func NewRejected(p Postgres) Rejected {
	return Rejected{
		record:   RejectedRecord{},
		postgres: p,
	}
}

func (e Rejected) Select(b io.Reader) ([]byte, error) {

	fmt.Println("Select not emplyment yet")

	return nil, nil
}

func (e Rejected) Update(b io.Reader) error {

	if err := e.postgres.Client.Where("id_i=?", e.record.IDI).Omit("country", "city").Assign(e).FirstOrCreate(&e).Error; err != nil {
		return err
	}
	return nil
}

func (e Rejected) Delete(b io.Reader) (error) {

	fmt.Println("Delete not emplyment yet")

	return nil
}

func (e Rejected)readBody(b io.Reader) (*RejectedRecord, error) {

	record := new(RejectedRecord)

	body, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, record); err != nil {
		return nil, err
	}

	return record, nil

}
