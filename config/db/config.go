package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alinabylkova/seen-task/config/env"
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
