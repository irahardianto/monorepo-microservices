package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/irahardianto/monorepo-microservices/auth/httphandler/interactor"
	"github.com/irahardianto/monorepo-microservices/auth/storage/mongodb"
)

type AuthHandler struct {
	login   interactor.LoginInteractor
	storage *mongodb.Storage
}

func NewAuthHandler(login interactor.LoginInteractor,
	storage *mongodb.Storage) *AuthHandler {
	return &AuthHandler{
		login:   login,
		storage: storage,
	}
}

// Login is a handler for hhtp POST - "/login"
// return auth token with refresh token
func (ah *AuthHandler) Login() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResourse UserResource
		// Decode the incoming user json
		err := json.NewDecoder(r.Body).Decode(&dataResourse)
		if err != nil {
			panic(err)
		}

		token, err := ah.login.Login(dataResourse.Data)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}

		setAuthAndRefreshCookies(&w, token.AuthToken, token.RefreshToken)
		w.Header().Set("X-CSRF-Token", token.CSRFKey)
		w.WriteHeader(http.StatusOK)
	})
}

func setAuthAndRefreshCookies(w *http.ResponseWriter, authTokenString string, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &refreshCookie)
}

func (ah *AuthHandler) GetReadiness() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ah.storage.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		}
	})
}

func (ah *AuthHandler) GetLiveness() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
