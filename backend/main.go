package main

import (
	"github.com/coma64/bahn-alarm-backend/cmd"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	setupLogging()
	cmd.Execute()
}

func setupLogging() {
	if config.Conf.Debug {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		zerolog.DurationFieldUnit = time.Minute
	}

	log.Logger = log.Logger.Level(zerolog.Level(config.Conf.LogLevel))
}
