package authenticator

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/irahardianto/monorepo-microservices/auth/model"
	"github.com/irahardianto/monorepo-microservices/package/hasher"
)

const (
	authExpTime         = 15 * time.Minute
	refreshTokenExpTIme = 72 * time.Hour
)

func CreateNewToken(uuid, username string, key []byte) (model.Token, error) {
	var token model.Token
	var err error

	token.CSRFKey, err = GenerateRandomString(32)
	if err != nil {
		return token, err
	}

	subject := fmt.Sprintf("%s-%s", uuid, username)
	authSubject := hasher.SHA256(subject)

	token.AuthToken, err = generateAuthToken(authSubject, token.CSRFKey, key)
	if err != nil {
		return token, err
	}

	token.RefreshToken, err = generateRefreshToken(authSubject, uuid, token.CSRFKey, key)
	if err != nil {
		return token, err
	}

	return token, nil
}

// GenerateAuthToken generate auth token with HS256 methos
// included standard claim
// subject is combination from UUID and username with hash
func generateAuthToken(subject, csrfSecret string, key []byte) (string, error) {
	authTokenExp := time.Now().Add(authExpTime).Unix()

	authClaim := model.TokenClaim{
		jwt.StandardClaims{
			Subject:   subject,
			ExpiresAt: authTokenExp,
		},
		csrfSecret,
	}

	authJwt := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), authClaim)
	return authJwt.SignedString(key)
}

func generateRefreshToken(subject, uuid, csrfSecret string, key []byte) (string, error) {
	authTokenExp := time.Now().Add(refreshTokenExpTIme).Unix()

	refreshClaims := model.TokenClaim{
		jwt.StandardClaims{
			Id:        uuid,
			Subject:   subject,
			ExpiresAt: authTokenExp,
		},
		csrfSecret,
	}

	// create a signer for rsa 256
	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshClaims)

	// generate the refresh token string
	return refreshJwt.SignedString(key)
}
