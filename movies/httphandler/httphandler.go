package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/irahardianto/monorepo-microservices/movies/storage"

	mgo "gopkg.in/mgo.v2"
)

// Handler for HTTP Get - "/movies"
// Returns all Movie documents
func GetMovies(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		movies := s.GetAll()
		j, err := json.Marshal(MoviesResource{Data: movies})
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// Handler for HTTP Post - "/movies"
// Insert a new Movie document
func CreateMovie(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResourse MovieResource
		// Decode the incoming Mov08ie json
		err := json.NewDecoder(r.Body).Decode(&dataResourse)
		if err != nil {
			panic(err)
		}

		movie := &dataResourse.Data

		s.Create(movie)
		j, err := json.Marshal(dataResourse)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// Handler for HTTP Get - "/movies/{id}"
// Get movie by id
func GetMovieById(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// Get movie by id
		movie, err := s.GetById(id)
		if err != nil {
			if err == mgo.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				panic(err)
			}
		}

		j, err := json.Marshal(MovieResource{Data: movie})
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

// Handler for HTTP Delete - "/movies/{id}"
// Delete movie by id
func DeleteMovie(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		err := s.Delete(id)
		if err != nil {
			if err == mgo.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				panic(err)
			}
		}
		w.WriteHeader(http.StatusNoContent)
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
