package main

import (
	"github.com/coma64/bahn-alarm-backend/db"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"os"
)

// See goose.run for available sub commands
func main() {
	dbConn, err := goose.OpenDBWithDriver("postgres", db.CreateConnectionStr())
	if err != nil {
		log.Err(err).Msg("failed to open db")
	}

	defer func() {
		if err = dbConn.Close(); err != nil {
			log.Err(err).Msg("failed to close db")
		}
	}()

	if err = goose.Run(os.Args[1], dbConn, "migrations", os.Args[2:]...); err != nil {
		log.Err(err).Send()
	}
}
