package httphandler

import "github.com/irahardianto/monorepo-mocroservices/movies/model"

type (
	// For Get - /movies
	MoviesResource struct {
		Data []model.Movie `json:"data"`
	}
	// For Post/Put - /movies
	MovieResource struct {
		Data model.Movie `json:"data"`
	}
)
