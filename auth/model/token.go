package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type TokenClaim struct {
	jwt.StandardClaims
	Csrf string `json:"csrf"`
}

type Token struct {
	AuthToken    string
	RefreshToken string
	CSRFKey      string
}
