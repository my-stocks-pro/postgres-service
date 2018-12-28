package postgres

import (
	"github.com/jinzhu/gorm"
	"io"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type EarningsRecord struct {
	gorm.Model
	Timestamp int64  `gorm:"type:int" json:"timestamp"`
	Date      string `gorm:"size:100" json:"date"`
	IDI       int    `gorm:"size:100" json:"idi"`
	Download  int    `gorm:"size:100" json:"download"`
	Category  string `gorm:"size:100" json:"category"`
	Country   string `gorm:"size:100" json:"country"`
	City      string `gorm:"size:100" json:"city"`
}

type Earnings struct {
	record   EarningsRecord
	postgres Postgres
}

func NewEarnings(p Postgres) Earnings {
	return Earnings{
		record:   EarningsRecord{},
		postgres: p,
	}
}

func (e Earnings) Select(b io.Reader) ([]byte, error) {
	//record, err := e.readBody(b)
	//if err != nil {
	//	return nil, err
	//}
	//
	//rows, err := e.postgres.Client.Find(&EarningsRecord{}).Rows()
	//if err != nil {
	//	return nil, err
	//}
	//
	//row := new(EarningsRecord)
	//defer rows.Close()
	//for rows.Next() {
	//	rows.Scan((*EarningsRecord).ID)
	//
	//}

	fmt.Println("Select not emplyment yet")

	return nil, nil
}

func (e Earnings) Update(b io.Reader) error {
	record, err := e.readBody(b)
	if err != nil {
		return err
	}

	if err := e.postgres.Client.Where("id_i=?", record.IDI).
		Omit("country", "city").Assign(e).FirstOrCreate(&record).Error; err != nil {
		return err
	}
	return nil
}

func (e Earnings) Delete(b io.Reader) (error) {
	//record, err := e.readBody(b)
	//if err != nil {
	//	return err
	//}

	fmt.Println("Delete not emplyment yet")

	return nil
}

func (e Earnings)readBody(b io.Reader) (*EarningsRecord, error) {

	record := new(EarningsRecord)

	body, err := ioutil.ReadAll(b)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, record); err != nil {
		return nil, err
	}

	return record, nil

}
