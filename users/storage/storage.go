package storage

import "github.com/irahardianto/monorepo-mocroservices/users/model"

type Storage interface {
	GetAll() []model.User
	Create(user *model.User) error
	Delete(id string) error
	Ping() error
}
