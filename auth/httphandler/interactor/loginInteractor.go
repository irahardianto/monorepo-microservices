package interactor

import (
	"github.com/irahardianto/monorepo-microservices/auth/model"
)

type LoginInteractor interface {
	Login(user model.User) (model.Token, error)
}
