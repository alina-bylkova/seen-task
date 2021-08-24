package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/alinabylkova/seen-task/config/env"
	"github.com/alinabylkova/seen-task/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Instance is a database client instance
type Instance struct {
	postgreDb *sql.DB
	Gorm      *gorm.DB
}

// New creates and initializes DB config
func New(c *env.Config) (*Instance, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.DbHost,
		c.DbPort,
		c.DbUser,
		c.DbPassword,
		c.DbName,
	)
	gormDb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to connect database: %v", err)
		return nil, err
	}

	postgreDb, err := gormDb.DB()
	if err != nil {
		log.Printf("Failed to get db from GORM: %v", err)
		return nil, err
	}
	postgreDb.SetMaxIdleConns(1)
	postgreDb.SetMaxOpenConns(c.MaxConnectionPool)
	postgreDb.SetConnMaxLifetime(c.MaxConnectionTimeout)

	instance := &Instance{postgreDb: postgreDb, Gorm: gormDb}
	return instance, nil
}

// CloseDb closes database connection pool
func (i *Instance) CloseDb() {
	if err := i.postgreDb.Close(); err != nil {
		log.Printf("DB connection close failed: %v", err)
	}
}

func (i *Instance) Get(r *model.Recipient) ([]*model.Recipient, error) {
	recipients := []*model.Recipient{}
	result := i.Gorm.
		Where(r).
		Find(&recipients)
	if result.RowsAffected == 0 {
		if result.Error != nil {
			log.Print("Database error: ", result.Error)
			return nil, result.Error
		}
		log.Print("Recipient(s) not found: ", r)
		return nil, errors.New("Recipient(s) not found")
	}
	log.Printf("Get returned %d recipient(s)", result.RowsAffected)
	return recipients, nil
}

func (i *Instance) GetAll() ([]*model.Recipient, error) {
	recipients := []*model.Recipient{}
	result := i.Gorm.Find(&recipients)
	if result.RowsAffected == 0 {
		if result.Error != nil {
			log.Print("Database error: ", result.Error)
			return nil, result.Error
		}
		log.Print("Recipient table is empty")
		return nil, errors.New("Recipient table is empty")
	}
	log.Printf("Get returned %d recipients", result.RowsAffected)
	return nil, nil
}

func (i *Instance) Add(r *model.Recipient) (int64, error) {
	result := i.Gorm.Create(r)
	if result.RowsAffected == 0 {
		if result.Error != nil {
			log.Print("Database error: ", result.Error)
			return 0, result.Error
		}
		log.Print("Recipient wasn't stored to db")
		return 0, errors.New("Recipient wasn't stored to db")
	}
	return r.ID, nil
}
