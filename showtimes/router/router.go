package router

import (
	"github.com/go-chi/chi"

	"github.com/irahardianto/microservice-monorepo/showtimes/httphandler"
	"github.com/irahardianto/microservice-monorepo/showtimes/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/showtimes", func(r chi.Router) {
		r.Get("/", httphandler.GetShowTimes(s))
		r.Post("/", httphandler.CreateShowTime(s))
		r.Get("/{date}", httphandler.GetShowTimeByDate(s))
		r.Delete("/{id}", httphandler.DeleteShowTime(s))
	})

	return r
}
