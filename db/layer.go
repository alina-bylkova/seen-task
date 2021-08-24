package db

import (
	"github.com/alinabylkova/seen-task/model"
)

type Layer interface {
	Get(*model.Recipient) ([]*model.Recipient, error)
	GetAll() ([]*model.Recipient, error)
	Add(*model.Recipient) (int64, error)
}
