package serve

import (
	"context"
	"errors"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/handlers"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/coma64/bahn-alarm-backend/watcher"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net/http"
)

var Cmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start the API server",
	Long:    "Starts the main server, the metrics server and the bahn API watcher",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	go watcher.WatchBahnApi(context.Background())
	go startPrometheusMetricsServer()

	e := createEchoServer()
	e.Use(middlewares...)

	bahnAlarmApi := &handlers.BahnAlarmApi{}
	server.RegisterHandlers(e, bahnAlarmApi)

	if err := e.Start(config.Conf.Bind); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("Failed to start bahn alarm server")
	}
}

func startPrometheusMetricsServer() {
	metricsServer := createEchoServer()
	metricsServer.HideBanner = true
	metricsServer.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogURI:    true,
			LogStatus: true,
			LogError:  true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				log.Info().
					Str("server", "metrics").
					Str("URI", v.URI).
					Int("status", v.Status).
					Dur("duration", v.Latency).
					Msg("metrics request")

				return nil
			},
		}),
	)

	metricsServer.GET("/metrics", echoprometheus.NewHandler())
	if err := metricsServer.Start(":2112"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("Failed to start metrics server")
	}
}

func createEchoServer() *echo.Echo {
	e := echo.New()
	e.Debug = config.Conf.Debug
	e.HideBanner = !config.Conf.Debug
	e.HidePort = !config.Conf.Debug
	return e
}
