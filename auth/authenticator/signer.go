package authenticator

import (
	"fmt"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
)

// Sign will generate JWT token with expiry time
func Sign(key []byte) (string, error) {
	claims := jws.Claims{}
	claims.Set("user", "hellouser")
	claims.SetExpiration(time.Now().Add(5 * time.Minute))

	token := jws.NewJWT(claims, crypto.SigningMethodHS256)
	serializedToken, err := token.Serialize(key)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(serializedToken), nil
}

// Validate will validate JWT token and return error if invalid
func Validate(token, key []byte) error {
	newToken, err := jws.ParseJWT([]byte(string(token)))
	if err != nil {
		return err
	}

	err = newToken.Validate(key, crypto.SigningMethodHS256)
	return err
}
