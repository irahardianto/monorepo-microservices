package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/irahardianto/monorepo-microservices/auth/model"

	"github.com/spf13/viper"

	"github.com/irahardianto/monorepo-microservices/auth/storage"
	"github.com/irahardianto/monorepo-microservices/package/authenticator/jwt"
)

// Authenticate is a Handler for HTTP POST - "/authenticate"
// Returns token for valid request
func Authenticate(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResourse UserResource
		// Decode the incoming Mov08ie json
		err := json.NewDecoder(r.Body).Decode(&dataResourse)
		if err != nil {
			panic(err)
		}

		usrData := &dataResourse.Data

		user, _ := s.GetByUsernameAndPassword(usrData.Username, usrData.Password)
		if user.ID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint("error: unauthorized user")))
		}

		key := viper.GetString("app.jwt_key")
		token, err := jwt.Sign([]byte(key))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(fmt.Sprint("error: unauthorized user")))
		}

		result, err := json.Marshal(model.AuthReponse{
			ID:       user.ID,
			Username: user.Username,
			Token:    token,
		})
		if err != nil {
			panic(err)
		}

		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	})
}

func GetReadiness(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := s.Ping()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", err)))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		}
	})
}

func GetLiveness() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
