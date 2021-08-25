package db

import (
	"github.com/alinabylkova/seen-task/model"
)

// Layer is an interface that is used to create and read records from database
type Layer interface {
	Get(*model.Recipient) ([]*model.Recipient, error)
	GetAll() ([]*model.Recipient, error)
	Add(*model.Recipient) (int64, error)
}
