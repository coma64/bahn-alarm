package main

import (
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/handlers"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

var authWhitelist = []string{
	"/auth/login",
	"/auth/register",
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

func main() {
	e := echo.New()
	e.Debug = config.Conf.Debug

	if config.Conf.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.DurationFieldUnit = time.Minute
	}

	e.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogError:  true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				var event *zerolog.Event
				status := v.Status
				if _, isHttpErr := v.Error.(*echo.HTTPError); v.Error != nil && !isHttpErr {
					event = log.Err(v.Error)
					status = http.StatusInternalServerError
				} else {
					event = log.Info()
				}

				event.
					Str("URI", v.URI).
					Int("status", status).
					Msg("request")

				return nil
			},
		}),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"http://localhost:5000"},
			AllowCredentials: true,
		}),
		middleware.TimeoutWithConfig(middleware.TimeoutConfig{
			Timeout: time.Duration(config.Conf.RequestTimeoutSeconds) * time.Second,
		}),
		echojwt.WithConfig(echojwt.Config{
			SigningKey:     []byte(config.Conf.Jwt.Secret),
			SigningMethod:  echojwt.AlgorithmHS256,
			TokenLookup:    "cookie:" + config.Conf.Jwt.Cookie,
			SuccessHandler: setUsernameFromJwtToken,
			Skipper: func(c echo.Context) bool {
				path := c.Request().URL.Path
				for _, whitelistedPath := range authWhitelist {
					if path == whitelistedPath {
						return true
					}
				}
				return false
			},
		}),
	)

	bahnAlarmApi := &handlers.BahnAlarmApi{}
	server.RegisterHandlers(e, bahnAlarmApi)

	e.File("/docs", "static/docs.html")
	e.Static("/static/swagger", "swagger-ui/dist")
	e.File("/openapi.yml", "openapi.yml")

	log.Fatal().Err(e.Start(config.Conf.Bind))
}
