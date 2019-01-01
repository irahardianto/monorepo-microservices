package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/irahardianto/monorepo-microservices/showtimes/storage"
	mgo "gopkg.in/mgo.v2"
)

func GetShowTimes(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		showtimes := s.GetAll()
		j, err := json.Marshal(ShowTimesResource{Data: showtimes})
		if err != nil {
			panic(err)
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

func CreateShowTime(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var dataResource ShowTimeResource
		// Decode the incoming ShowTime json
		err := json.NewDecoder(r.Body).Decode(&dataResource)
		if err != nil {
			panic(err)
		}
		showtime := &dataResource.Data

		s.Create(showtime)
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

func GetShowTimeByDate(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		date := chi.URLParam(r, "date")

		// Get showtime by date
		showtime, err := s.GetByDate(date)
		if err != nil {
			panic(err)
		}
		// Create data for the response
		j, err := json.Marshal(ShowTimeResource{Data: showtime})
		if err != nil {
			panic(err)
		}
		// Send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	})
}

func DeleteShowTime(s storage.Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		// Remove showtime by id
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
