package main

import (
	"github.com/coma64/bahn-alarm-backend/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	db, err := goose.OpenDBWithDriver("postgres", config.Conf.Db.Dsn)
	if err != nil {
		log.Err(err).Msg("failed to open db")
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Err(err).Msg("failed to close db")
		}
	}()

	if err = goose.Run(os.Args[1], db, "migrations", os.Args[2:]...); err != nil {
		log.Err(err).Send()
	}
}
