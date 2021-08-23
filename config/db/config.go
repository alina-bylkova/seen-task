package dbconfig

import (
	"fmt"

	"github.com/alinabylkova/seen-task/config/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New creates and initializes DB config
func New(c *env.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DbHost, c.DbPort, c.DbUser, c.DbPassword, c.DbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect database: %v", err)
	}

	postgreDb, err := db.DB()
	if err != nil {
		fmt.Printf("Failed to get db from GORM: %v", err)
	}
	postgreDb.SetMaxIdleConns(1)
	postgreDb.SetMaxOpenConns(c.MaxConnectionPool)
	postgreDb.SetConnMaxLifetime(c.MaxConnectionTimeout)

	return db
}
