package main

import (
	"context"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/handlers"
	"github.com/coma64/bahn-alarm-backend/metrics"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/coma64/bahn-alarm-backend/watcher"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"strings"
	"time"
)

var authWhitelist = []string{
	"/auth/login",
	"/auth/register",
	"/docs",
	"/static/swagger",
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
	e.HideBanner = !config.Conf.Debug
	e.HidePort = !config.Conf.Debug

	if config.Conf.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.DurationFieldUnit = time.Minute
	}

	log.Logger = log.Logger.Level(zerolog.Level(config.Conf.LogLevel))

	watcherCtx, cancelWatcher := context.WithCancel(context.Background())
	defer cancelWatcher()
	go func() {
		watcher.WatchBahnApi(watcherCtx)
	}()

	go exposePrometheusMetrics()

	e.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogError:  true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				metrics.RequestDuration.Observe(v.Latency.Seconds())

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
	)

	bahnAlarmApi := &handlers.BahnAlarmApi{}
	server.RegisterHandlers(e, bahnAlarmApi)

	e.File("/docs", "static/docs.html")
	e.File("/docs/openapi.yml", "openapi.yml")
	e.Static("/static/swagger", "swagger-ui/dist")

	log.Fatal().Err(e.Start(config.Conf.Bind)).Send()
}

func exposePrometheusMetrics() {
	handler := http.NewServeMux()
	handler.Handle("/metrics", promhttp.Handler())
	log.Fatal().Err(http.ListenAndServe(":2112", handler)).Send()
}
