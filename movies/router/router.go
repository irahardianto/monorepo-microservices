package router

import (
	"github.com/go-chi/chi"

	"github.com/irahardianto/microservice-monorepo/movies/httphandler"
	"github.com/irahardianto/microservice-monorepo/movies/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/movies", func(r chi.Router) {
		r.Get("/", httphandler.GetMovies(s))
		r.Post("/", httphandler.CreateMovie(s))
		r.Get("/{id}", httphandler.GetMovieById(s))
		r.Delete("/{id}", httphandler.DeleteMovie(s))
	})

	return r
}
