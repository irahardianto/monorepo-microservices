package authenticator

import (
	"errors"
	"time"

	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"github.com/spf13/viper"
)

func ValidateToken(token []byte) error {
	key := viper.GetString("app.jwt_key")
	jwtToken, err := jws.ParseJWT([]byte(string(token)))
	if err != nil {
		return err
	}

	exp := time.Minute
	fn := func(c jws.Claims) error {
		if c.Get("user") != "hellouserrr" {
			return errors.New("invalid user")
		}
		return nil
	}
	v := jws.NewValidator(jws.Claims{}, exp, exp, fn)

	err = jwtToken.Validate([]byte(key), crypto.SigningMethodHS256, v)
	return err
}
