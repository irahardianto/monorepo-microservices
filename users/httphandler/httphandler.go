package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/irahardianto/microservice-monorepo/users/storage"
	mgo "gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/users"
// Returns all User documents
func GetUsers(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := s.GetAll()
		j, err := json.Marshal(UsersResource{Data: users})
		if err != nil {
			panic(err)
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// Handler for HTTP Post - "/users"
// Create a new Showtime document
func CreateUser(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResource UserResource
		// Decode the incoming User json
		err := json.NewDecoder(r.Body).Decode(&dataResource)
		if err != nil {
			panic(err)
		}
		user := &dataResource.Data
		// Create User
		s.Create(user)
		// Create response data
		j, err := json.Marshal(dataResource)
		if err != nil {
			panic(err)
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// Handler for HTTP Delete - "/users/{id}"
// Delete a User document by id
func DeleteUser(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// Remove user by id
		err := s.Delete(id)
		if err != nil {
			if err == mgo.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				panic(err)
			}
		}

		// Send response back
		w.WriteHeader(http.StatusNoContent)
	})
}
