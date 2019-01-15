package httphandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/irahardianto/monorepo-microservices/auth/httphandler/interactor"
	"github.com/irahardianto/monorepo-microservices/auth/storage/mongodb"
)

type AuthHandler struct {
	authInteractor interactor.AuthenticationInteractor
	storage        *mongodb.Storage
}

func NewAuthHandler(authInteractor interactor.AuthenticationInteractor,
	storage *mongodb.Storage) *AuthHandler {
	return &AuthHandler{
		authInteractor: authInteractor,
		storage:        storage,
	}
}

// Login is a handler for http POST - "/login"
// return auth token with refresh token
func (ah *AuthHandler) Login() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResourse UserResource
		// Decode the incoming user json
		err := json.NewDecoder(r.Body).Decode(&dataResourse)
		if err != nil {
			panic(err)
		}

		token, err := ah.authInteractor.Login(dataResourse.Data)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
		}

		setAuthCookie(&w, token.AuthToken, token.RefreshToken)
		setRefreshCookie(&w, token.AuthToken, token.RefreshToken)

		w.Header().Set("X-CSRF-Token", token.CSRFKey)
		w.WriteHeader(http.StatusOK)
	})
}

// Authentication is a handler for http POST - "/authentication"
// return HTTP ok if token valid and HTTP 400 if not valid
func (ah *AuthHandler) Authentication() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie("AuthToken")
		if err == http.ErrNoCookie {
			ah.nullifyTokenCookies(&w, r)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		} else if err != nil {
			ah.nullifyTokenCookies(&w, r)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		refreshToken, err := r.Cookie("RefreshToken")
		if err == http.ErrNoCookie {
			ah.nullifyTokenCookies(&w, r)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		} else if err != nil {
			ah.nullifyTokenCookies(&w, r)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		csrfKey := getCSRFKey(r)

		token, err := ah.authInteractor.Authenticate(authCookie.Value, refreshToken.Value, csrfKey)
		if err != nil {
			println(err.Error())
			if err.Error() == "Unauthorized" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			} else {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}

		// @adam-hanna: Change this. Only allow whitelisted origins! Also check referer header

		setAuthCookie(&w, token.AuthToken, token.RefreshToken)
		setRefreshCookie(&w, token.AuthToken, token.RefreshToken)
		w.Header().Set("X-CSRF-Token", token.CSRFKey)
		w.WriteHeader(http.StatusOK)
	})
}

func getCSRFKey(r *http.Request) string {
	csrfFromFrom := r.FormValue("X-CSRF-Token")

	if csrfFromFrom != "" {
		return csrfFromFrom
	} else {
		return r.Header.Get("X-CSRF-Token")
	}
}

func setAuthCookie(w *http.ResponseWriter, authTokenString string, refreshTokenString string) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    authTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &authCookie)
}

func setRefreshCookie(w *http.ResponseWriter, authTokenString string, refreshTokenString string) {
	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenString,
		HttpOnly: true,
	}

	http.SetCookie(*w, &refreshCookie)
}

func (ah *AuthHandler) nullifyTokenCookies(w *http.ResponseWriter, r *http.Request) {
	authCookie := http.Cookie{
		Name:     "AuthToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(*w, &authCookie)

	refreshCookie := http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		Expires:  time.Now().Add(-1000 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(*w, &refreshCookie)

	// if present, revoke the refresh cookie from our db
	RefreshCookie, refreshErr := r.Cookie("RefreshToken")
	if refreshErr == http.ErrNoCookie {
		return
	} else if refreshErr != nil {
		log.Panic("panic: %+v", refreshErr)
		http.Error(*w, http.StatusText(500), 500)
	}

	ah.authInteractor.RevokeRefreshToken(RefreshCookie.Value)
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
