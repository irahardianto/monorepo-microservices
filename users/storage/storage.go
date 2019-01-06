package storage

import (
	"github.com/irahardianto/monorepo-microservices/users/model"
)

type Storage interface {
	GetAll() []model.User
	Create(user *model.User) error
	Delete(id string) error
	Ping() error
}
