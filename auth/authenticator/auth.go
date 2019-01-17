package authenticator

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/irahardianto/monorepo-microservices/auth/usecase/storage"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/irahardianto/monorepo-microservices/auth/model"
	"github.com/irahardianto/monorepo-microservices/package/hasher"
)

const (
	authExpTime         = 10 * time.Second
	refreshTokenExpTIme = 72 * time.Hour
)

type Authenticator struct {
	authStorage storage.UserInteractor
}

func NewAuthenticator(authStorage storage.UserInteractor, key []byte) *Authenticator {
	return &Authenticator{
		authStorage: authStorage,
	}
}

// CreateNewToken generate new token with JWT using HS256 method
// return token and error
func (auth *Authenticator) CreateNewToken(uuid, username string, key []byte) (model.Token, error) {
	var token model.Token
	var err error

	token.CSRFKey, err = GenerateRandomString(32)
	if err != nil {
		return token, err
	}

	subject := fmt.Sprintf("%s-%s", uuid, username)
	authSubject := hasher.SHA256(subject)

	token.AuthToken, err = auth.generateAuthToken(authSubject, token.CSRFKey, key)
	if err != nil {
		return token, err
	}

	token.RefreshToken, err = auth.generateRefreshToken(authSubject, uuid, token.CSRFKey, key)
	if err != nil {
		return token, err
	}

	return token, nil
}

// GenerateAuthToken generate auth token with HS256 methos
// included standard claim
// subject is combination from UUID and username with hash
func (auth *Authenticator) generateAuthToken(subject, csrfSecret string, key []byte) (string, error) {
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

func (auth *Authenticator) generateRefreshToken(tokenID, subject, csrfSecret string, key []byte) (string, error) {
	authTokenExp := time.Now().Add(refreshTokenExpTIme).Unix()

	auth.authStorage.StoreRefreshToken(model.RefreshToken{Token: tokenID})

	refreshClaims := model.TokenClaim{
		jwt.StandardClaims{
			Id:        tokenID,
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

func (auth *Authenticator) getTokenWithClaim(token string, key []byte) (*jwt.Token, *model.TokenClaim, error) {
	refreshToken, err := jwt.ParseWithClaims(token, &model.TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	tokenClaim, ok := refreshToken.Claims.(*model.TokenClaim)
	if !ok {
		return nil, nil, err
	}

	return refreshToken, tokenClaim, err
}

func (auth *Authenticator) updateRefreshToken(oldRefreshToken string, key []byte) (newRefreshToken string, err error) {
	_, oldRefreshTokenClaims, err := auth.getTokenWithClaim(oldRefreshToken, key)

	refreshTokenExp := time.Now().Add(refreshTokenExpTIme).Unix()

	refreshClaims := model.TokenClaim{
		jwt.StandardClaims{
			Id:        oldRefreshTokenClaims.StandardClaims.Id, // jti
			Subject:   oldRefreshTokenClaims.StandardClaims.Subject,
			ExpiresAt: refreshTokenExp,
		},
		oldRefreshTokenClaims.Csrf,
	}

	// create a signer for rsa 256
	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshClaims)

	// generate the refresh token string
	return refreshJwt.SignedString(key)
}

func (auth *Authenticator) updateAuthToken(rawRefreshToken, rawAuthToken string, key []byte) (model.Token, error) {
	var newToken model.Token

	refreshToken, refreshTokenClaims, err := auth.getTokenWithClaim(rawRefreshToken, key)
	storedToken, err := auth.authStorage.GetRefreshToken(refreshTokenClaims.StandardClaims.Id)
	if err != nil {
		return newToken, err
	}

	if storedToken.ID != "" {
		if refreshToken.Valid {
			_, authTokenClaim, _ := auth.getTokenWithClaim(rawAuthToken, key)
			newToken.CSRFKey, err = GenerateRandomString(32)
			if err != nil {
				return newToken, err
			}

			newToken.AuthToken, err = auth.generateAuthToken(authTokenClaim.StandardClaims.Subject, newToken.CSRFKey, key)
			return newToken, err
		}
		// the refresh token has expired!
		// Revoke the token in our db and require the user to login again
		log.Println("Refresh token has expired!")
		auth.authStorage.DeleteRefreshToken(refreshTokenClaims.StandardClaims.Id)

		err = errors.New("Unauthorized")
		return newToken, err
	}

	log.Println("Refresh token has been revoked!")
	err = errors.New("Unauthorized")
	return newToken, err
}

func (auth *Authenticator) updateRefreshTokenCsrf(oldRefreshTokenString string, newCsrfString string, key []byte) (string, error) {
	_, oldRefreshTokenClaims, err := auth.getTokenWithClaim(oldRefreshTokenString, key)
	if err != nil {
		return "", err
	}

	refreshClaims := model.TokenClaim{
		jwt.StandardClaims{
			Id:        oldRefreshTokenClaims.StandardClaims.Id, // jti
			Subject:   oldRefreshTokenClaims.StandardClaims.Subject,
			ExpiresAt: oldRefreshTokenClaims.StandardClaims.ExpiresAt,
		},
		newCsrfString,
	}

	// create a signer for rsa 256
	refreshJwt := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), refreshClaims)

	// generate the refresh token string
	return refreshJwt.SignedString(key)
}

// AuthenticateToken check token and return new token
func (auth *Authenticator) AuthenticateToken(authToken, refreshToken, csrfToken string, key []byte) (model.Token, error) {
	var token model.Token
	tokenWithClaim, authTokenClaims, err := auth.getTokenWithClaim(authToken, key)

	if csrfToken != authTokenClaims.Csrf {
		log.Println("CSRF token doesn't match jwt!")
		return token, errors.New("Unauthorized")
	}

	if tokenWithClaim.Valid {
		token.CSRFKey = authTokenClaims.Csrf
		token.AuthToken = authToken
		token.RefreshToken, err = auth.updateRefreshToken(refreshToken, key)
		return token, err
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
			token, err = auth.updateAuthToken(refreshToken, authToken, key)
			if err != nil {
				return token, err
			}

			// update the exp of refresh token string
			token.RefreshToken, err = auth.updateRefreshToken(refreshToken, key)
			if err != nil {
				return token, err
			}

			// update the csrf string of the refresh token
			token.RefreshToken, err = auth.updateRefreshTokenCsrf(token.RefreshToken, token.CSRFKey, key)
			return token, err
		}
		log.Println("Error in auth token")
		return token, errors.New("Error in auth token")

	} else {
		return token, errors.New("Error in auth token")
	}
}

// RevokeRefreshToken remove refresh token from storage
func (auth *Authenticator) RevokeRefreshToken(refreshToken string) error {
	return auth.authStorage.DeleteRefreshToken(refreshToken)
}
