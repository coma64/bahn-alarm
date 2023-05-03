package auth

import (
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJwt(username string) (string, error) {
	expiration := time.Now().Add(time.Hour * 24 * time.Duration(config.Conf.Jwt.ExpirationDays))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": expiration.Unix(),
	})

	return token.SignedString([]byte(config.Conf.Jwt.Secret))
}
