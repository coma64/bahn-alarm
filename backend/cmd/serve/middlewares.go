package serve

import (
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/metrics"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo-contrib/echoprometheus"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
	"time"
)

var authWhitelist = []string{
	"/auth/login",
	"/auth/register",
	"/docs",
	"/static/swagger",
}

var middlewares = []echo.MiddlewareFunc{
	echoprometheus.NewMiddleware(metrics.Prefix),
	middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogError:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			var event *zerolog.Event
			status := v.Status
			if httpErr, isHttpErr := v.Error.(*echo.HTTPError); v.Error != nil && !isHttpErr {
				event = log.Err(v.Error)
				status = http.StatusInternalServerError
			} else if v.Error != nil && isHttpErr && httpErr.Code >= 500 {
				event = log.Err(v.Error)
			} else {
				event = log.Info()
			}

			event.
				Str("server", "api").
				Str("URI", v.URI).
				Int("status", status).
				Dur("duration", v.Latency).
				Msg("request")

			return nil
		},
	}),
	middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     config.Conf.Requests.CorsOrigins,
		AllowCredentials: true,
	}),
	middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: time.Duration(config.Conf.Requests.TimeoutSeconds) * time.Second,
	}),
	echojwt.WithConfig(echojwt.Config{
		SigningKey:     []byte(config.Conf.Jwt.Secret),
		SigningMethod:  echojwt.AlgorithmHS256,
		TokenLookup:    "cookie:" + config.Conf.Jwt.Cookie,
		SuccessHandler: setUsernameFromJwtToken,
		Skipper: func(c echo.Context) bool {
			path := c.Request().URL.Path
			for _, whitelistedPath := range authWhitelist {
				if strings.HasPrefix(path, whitelistedPath) {
					return true
				}
			}
			return false
		},
	}),
}

func setUsernameFromJwtToken(ctx echo.Context) {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return
	}

	var claims jwt.MapClaims
	claims, ok = token.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	ctx.Set("username", claims["sub"].(string))
}
