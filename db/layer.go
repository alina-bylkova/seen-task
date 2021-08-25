package db

import (
	"github.com/alinabylkova/seen-task/model"
	"github.com/alinabylkova/seen-task/model/dto"
)

// Layer is an interface that is used to create and read records from database
type Layer interface {
	Get(*model.Recipient) ([]*model.Recipient, error)
	GetAll() ([]*model.Recipient, error)
	AddRecipient(*model.Recipient) (int64, error)
	AddEvent(*dto.Event) error
}
