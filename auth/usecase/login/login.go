package login

import (
	"fmt"

	"github.com/irahardianto/monorepo-microservices/auth/authenticator"
	"github.com/irahardianto/monorepo-microservices/auth/model"
	"github.com/irahardianto/monorepo-microservices/auth/usecase/storage"
	"github.com/irahardianto/monorepo-microservices/package/hasher"
	"github.com/spf13/viper"
)

type Login struct {
	storage storage.UserInteractor
	jwtKey  []byte
}

func NewLogin(storage storage.UserInteractor) *Login {
	strKey := viper.GetString("app.jwt_key")

	return &Login{
		storage: storage,
		jwtKey:  []byte(strKey),
	}
}

// Login check the username and password
// then generate user token
func (lgn *Login) Login(user model.User) (model.Token, error) {
	var token model.Token

	hashPassword := hasher.SHA256(user.Password)

	user, err := lgn.storage.GetByUsernameAndPassword(user.Username, hashPassword)
	if err != nil {
		return token, err
	}

	return authenticator.CreateNewToken(fmt.Sprint(user.ID), user.Username, lgn.jwtKey)
}
