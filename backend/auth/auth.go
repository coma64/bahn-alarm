package auth

import (
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func GenerateJwt(username string) (jwtToken string, expiresAt time.Time, err error) {
	expiresAt = time.Now().Add(time.Hour * 24 * time.Duration(config.Conf.Jwt.ExpirationDays))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": expiresAt.Unix(),
	})

	jwtToken, err = token.SignedString([]byte(config.Conf.Jwt.Secret))
	return jwtToken, expiresAt, err
}
