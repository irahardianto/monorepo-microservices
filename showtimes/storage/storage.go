package storage

import (
	"github.com/irahardianto/microservice-monorepo/showtimes/model"
)

type Storage interface {
	Create(showtime *model.ShowTime) error
	GetAll() []model.ShowTime
	GetByDate(date string) (showtime model.ShowTime, err error)
	Delete(id string) error
}
