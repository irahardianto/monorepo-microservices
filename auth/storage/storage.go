package storage

import "github.com/irahardianto/monorepo-microservices/auth/model"

type Storage interface {
	GetByUsernameAndPassword(username, password string) (model.User, error)
	Ping() error
}
