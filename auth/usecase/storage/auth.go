package storage

import (
	"github.com/irahardianto/monorepo-microservices/auth/model"
)

// UserInteractor is an interface to interaction with User module
type UserInteractor interface {
	GetByUsernameAndPassword(username, password string) (model.User, error)
	StoreRefreshToken(token model.RefreshToken) error
	DeleteRefreshToken(token string) error
	GetRefreshToken(token string) (model.RefreshToken, error)
	Ping() error
}
