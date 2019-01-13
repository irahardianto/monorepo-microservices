package router

import (
	"github.com/go-chi/chi"
	"github.com/irahardianto/monorepo-microservices/auth/httphandler"
	"github.com/irahardianto/monorepo-microservices/auth/storage/mongodb"
	"github.com/irahardianto/monorepo-microservices/auth/usecase/login"
)

func InitRouter(r *chi.Mux, s *mongodb.Storage) *chi.Mux {
	authHandler := initDependencies(s)
	r.Post("/login", authHandler.Login())
	r.Get("/healthy", authHandler.GetReadiness())
	r.Get("/healthz", authHandler.GetLiveness())

	return r
}

func initDependencies(storage *mongodb.Storage) httphandler.AuthHandler {
	login := login.NewLogin(storage)
	authHandler := httphandler.NewAuthHandler(login, storage)
	return *authHandler
}
