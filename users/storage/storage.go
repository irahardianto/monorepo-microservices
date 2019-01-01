package storage

import (
	"github.com/irahardianto/microservice-monorepo/users/model"
)

type Storage interface {
	GetAll() []model.User
	Create(user *model.User) error
	Delete(id string) error
}
