package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"github.com/my-stocks-pro/api-server/models"
)

type Postgres interface {
	//Ping() error
	//Select(client *gorm.DB) ([]byte, error)
	Update(client *gorm.DB) error
	//Delete() error
}

type PostgresType struct {
	config Config
	client *gorm.DB
}

func NewPostgres(config Config, client *gorm.DB) PostgresType {
	return PostgresType{
		config: config,
		client: client,
	}
}

func NewClient(c Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("sslmode=disable host=%s port=%s dbname=%s user=%s password=%s",
		c.PHOST, c.PPORT, c.PNAME, c.PUSER, c.PPASS)

	connection, err := gorm.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := connection.DB().Ping(); err != nil {
		return nil, err
	}

	if os.Getenv("MIGRATE") == "1" {
		connection.AutoMigrate(&models.Approve{}, &models.Earnings{})
		//connection.AutoMigrate()
	}

	return connection, nil
}

func (p PostgresType) Ping() error {
	if err := p.client.DB().Ping(); err != nil {
		return err
	}
	return nil
}

func (p PostgresType) Select() ([]byte, error) {

	p.client.Select().Where()

	return nil, nil
}

func (p PostgresType) Update() error {

	return nil
}

func (p PostgresType) Delete() error {

	return nil
}
