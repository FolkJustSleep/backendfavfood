package database

import (
	"database/sql"
	"fmt"
	"os"

	fiberlog "github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-template/data/model"
)

type PSQL struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}
type CRUD struct {
	DB *sql.DB
}

type ICRUD interface {
	Create(queryparams string, table string, args ...interface{}) error
	GetAll(table string) (*sql.Rows, error)
	GetAllByID(table string, id string) (*sql.Row, error)
	GetAllByEmpID(table string, emp_id string) (*sql.Row, error)
	GetFieldByID(table string, field string, id string) (*sql.Row, error)
	Update(table string, queryparams string, id string, args ...interface{}) (*sql.Row, error)
	Delete(table string, id string) error
}

func NewPSQL() *PSQL {
	return &PSQL{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		dbname:   os.Getenv("DB_NAME"),
	}
}

func (p *PSQL) ConnectGorm() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.host, p.port, p.user, p.password, p.dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}

	fiberlog.Info("Successfully connected to database")

	// Migrate the schema
	db.AutoMigrate(&model.User{}, &model.Logs{})
	db.AutoMigrate(&model.CashControl{})
	db.AutoMigrate(&model.Item{})
	db.AutoMigrate(&model.Delivery{})
	db.AutoMigrate(&model.Timestamp{})
	db.AutoMigrate(&model.Menu{})

	return db, nil
}


