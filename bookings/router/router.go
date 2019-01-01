package router

import (
	"github.com/go-chi/chi"

	"github.com/irahardianto/microservice-monorepo/bookings/httphandler"
	"github.com/irahardianto/microservice-monorepo/bookings/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/bookings", func(r chi.Router) {
		r.Get("/", httphandler.GetBookings(s))
		r.Post("/", httphandler.CreateBooking(s))
		r.Delete("/{id}", httphandler.DeleteBooking(s))
	})

	return r
}
