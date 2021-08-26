package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/alinabylkova/seen-task/config/env"
	"github.com/alinabylkova/seen-task/model"
	"github.com/alinabylkova/seen-task/model/dto"
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

// Get selects recipient(s) from the database based on the provided id, name, email and phone number
func (i *Instance) Get(r *model.Recipient) ([]*model.Recipient, error) {
	recipients := []*model.Recipient{}
	result := i.Gorm.
		Where(r).
		Find(&recipients)
	if err := checkDbResult(result, "Recipient(s) not found"); err != nil {
		return nil, err
	}
	log.Printf("Get returned %d recipient(s)", result.RowsAffected)
	return recipients, nil
}

// GetAll selects all recipients from the database
func (i *Instance) GetAll() ([]*model.Recipient, error) {
	recipients := []*model.Recipient{}
	result := i.Gorm.Find(&recipients)
	if err := checkDbResult(result, "Recipient table is empty"); err != nil {
		return nil, err
	}
	log.Printf("Get returned %d recipients", result.RowsAffected)
	return recipients, nil
}

// AddRecipient creates new recipient in the database based on provided name, email and phone number
func (i *Instance) AddRecipient(r *model.Recipient) (int64, error) {
	result := i.Gorm.Create(r)
	if err := checkDbResult(result, "Recipient wasn't stored to db"); err != nil {
		return 0, err
	}
	return r.ID, nil
}

// AddEvent creates or updates event in the database based on the provided recipient_id, video_id and event_type
func (i *Instance) AddEvent(e *dto.Event) error {
	event := &model.Event{}
	searchElement := &model.Event{
		RecipientID: e.RecipientID,
		VideoID:     e.VideoID,
	}

	searchResult := i.Gorm.Where(searchElement).Find(event)
	if searchResult.RowsAffected == 0 {
		if searchResult.Error != nil {
			log.Print("Database error: ", searchResult.Error)
			return &DbError{originalError: searchResult.Error}
		}
		event.RecipientID = e.RecipientID
		event.VideoID = e.VideoID

		updateEvent(e.EventType, event)

		createResult := i.Gorm.Create(event)
		return checkDbResult(createResult, "Event wasn't stored to db")
	}
	updateEvent(e.EventType, event)

	updateResult := i.Gorm.Save(event)
	return checkDbResult(updateResult, "Event wasn't stored to db")
}

func updateEvent(eventType string, event *model.Event) {
	if eventType == dto.LpHits {
		event.LpHits += 1
	}
	if eventType == dto.VideoPlays {
		event.VideoPlays += 1
	}
}

func checkDbResult(result *gorm.DB, message string) error {
	if result.RowsAffected == 0 {
		if result.Error != nil {
			log.Print("Database error: ", result.Error)
			return &DbError{originalError: result.Error}
		}
		log.Print(message)
		return errors.New(message)
	}
	return nil
}
