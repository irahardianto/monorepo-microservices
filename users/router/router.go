package router

import (
	"github.com/go-chi/chi"

	"github.com/irahardianto/monorepo-microservices/users/httphandler"
	"github.com/irahardianto/monorepo-microservices/users/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {

	r.Route("/users", func(r chi.Router) {
		r.Get("/", httphandler.GetUsers(s))
		r.Post("/", httphandler.CreateUser(s))
		r.Delete("/{id}", httphandler.DeleteUser(s))
	})

	return r
}
