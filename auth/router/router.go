package router

import (
	"github.com/go-chi/chi"
	"github.com/irahardianto/monorepo-microservices/auth/httphandler"
	"github.com/irahardianto/monorepo-microservices/auth/storage"
)

func InitRouter(r *chi.Mux, s storage.Storage) *chi.Mux {
	r.Post("/authenticate", httphandler.Authenticate(s))
	r.Get("/healthy", httphandler.GetReadiness(s))
	r.Get("/healthz", httphandler.GetLiveness())

	return r
}
