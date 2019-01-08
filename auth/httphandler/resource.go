package httphandler

import "github.com/irahardianto/monorepo-microservices/auth/model"

type (
	// For Post/Put - /users
	UserResource struct {
		Data model.User `json:"data"`
	}
)
