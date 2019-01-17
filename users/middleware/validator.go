package middleware

import (
	"fmt"
	"net/http"

	"github.com/irahardianto/monorepo-microservices/package/authenticator/jwt"
)

// ValidateToken is a middleware that validate token from header
func ValidateToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if err := jwt.ValidateToken([]byte(token)); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	return http.HandlerFunc(fn)
}
