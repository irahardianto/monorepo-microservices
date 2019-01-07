package httphandler

import "github.com/irahardianto/monorepo-microservices/showtimes/model"

type (
	// For Get - /showtimes
	ShowTimesResource struct {
		Data []model.ShowTime `json:"data"`
	}
	// For Post/Put - /showtimes
	ShowTimeResource struct {
		Data model.ShowTime `json:"data"`
	}
)
