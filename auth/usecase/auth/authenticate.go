package auth

import (
	"errors"

	"github.com/irahardianto/monorepo-microservices/auth/authenticator"
	"github.com/irahardianto/monorepo-microservices/auth/model"
	"github.com/irahardianto/monorepo-microservices/auth/usecase/storage"
	"github.com/irahardianto/monorepo-microservices/package/hasher"
	"github.com/spf13/viper"
)

type Authentication struct {
	storage       storage.UserInteractor
	authenticator *authenticator.Authenticator
	jwtKey        []byte
}

func NewAuthentication(storage storage.UserInteractor) *Authentication {
	strKey := viper.GetString("app.jwt_key")
	jwtKey := []byte(strKey)
	return &Authentication{
		storage:       storage,
		jwtKey:        jwtKey,
		authenticator: authenticator.NewAuthenticator(storage, jwtKey),
	}
}

// Login check the username and password
// then generate user token
func (auth *Authentication) Login(user model.User) (model.Token, error) {
	var token model.Token

	hashPassword := hasher.SHA256(user.Password)

	user, err := auth.storage.GetByUsernameAndPassword(user.Username, hashPassword)
	if err != nil {
		return token, err
	}

	return auth.authenticator.CreateNewToken(user.ID.Hex(), user.Username, auth.jwtKey)
}

// Authenticate check the auth,refresh and csrf token
// then refresh token expiry time
func (auth *Authentication) Authenticate(authToken, refreshToken, csrfToken string) (model.Token, error) {
	var token model.Token
	if csrfToken == "" {
		return token, errors.New("no csrf token")
	}

	return auth.authenticator.AuthenticateToken(authToken, refreshToken, csrfToken, auth.jwtKey)
}

// RevokeRefreshToken remove token from storage
func (auth *Authentication) RevokeRefreshToken(refreshToken string) error {
	return auth.authenticator.RevokeRefreshToken(refreshToken)
}
