package postgres

import (
	"github.com/jinzhu/gorm"
	"io"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type ApprovedRecord struct {
	gorm.Model
	Timestamp   int64  `gorm:"type:int" json:"timestamp"`
	IDI         string `gorm:"size:30" json:"id"`
	AddedDate   string `gorm:"size:30" json:"added_date"`
	Link        string `gorm:"size:1024" json:"link"`
	Description string `gorm:"size:10240" json:"description"`
}

type Approved struct {
	record   ApprovedRecord
	postgres Postgres
}

func NewApproved(p Postgres) Earnings {
	return Earnings{
		record:   EarningsRecord{},
		postgres: p,
	}
}

func (e Approved) Select(b io.Reader) ([]byte, error) {

	fmt.Println("Select not emplyment yet")

	return nil, nil
}

func (e Approved) Update(b io.Reader) error {
	if err := e.postgres.Client.Where("id_i=?", e.record.IDI).Omit("country", "city").Assign(e).FirstOrCreate(&e).Error; err != nil {
		return err
	}
	return nil
}

func (e Approved) Delete(b io.Reader) (error) {

	fmt.Println("Delete not emplyment yet")

	return nil
}

func (e Approved)readBody(b io.Reader) (*ApprovedRecord, error) {

	record := new(ApprovedRecord)

	body, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, record); err != nil {
		return nil, err
	}

	return record, nil

}
