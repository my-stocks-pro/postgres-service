package postgres

import (
	"fmt"
	"os"
	"io"

	"github.com/jinzhu/gorm"
	"github.com/my-stocks-pro/postgres-service/infrastructure"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Persist interface {
	Select(b io.Reader) ([]byte, error)
	Update(b io.Reader) error
	Delete(b io.Reader) error
}

type Postgres struct {
	Config infrastructure.Config
	Client *gorm.DB
}

func NewPostgres(config infrastructure.Config, client *gorm.DB) Postgres {
	return Postgres{
		Config: config,
		Client: client,
	}
}

func NewClient(c infrastructure.Config) (*gorm.DB, error) {
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
		connection.AutoMigrate(&ApprovedRecord{}, &EarningsRecord{}, &RejectedRecord{})
		//connection.AutoMigrate()
	}

	return connection, nil
}

func (p Postgres) Ping() error {
	if err := p.Client.DB().Ping(); err != nil {
		return err
	}
	return nil
}
