package cmd

import (
	"github.com/coma64/bahn-alarm-backend/db"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

// See goose.run for available sub commands
var migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "Runs goose with the passed in arguments",
	Long:    "Runs the goose command with a predefined connection. See goose.run for available sub commands",
	Args:    cobra.MinimumNArgs(1),
	Aliases: []string{"m"},
	Run: func(cmd *cobra.Command, args []string) {
		dbConn, err := goose.OpenDBWithDriver("postgres", db.CreateConnectionStr())
		if err != nil {
			log.Err(err).Msg("failed to open db")
		}

		defer func() {
			if err = dbConn.Close(); err != nil {
				log.Err(err).Msg("failed to close db")
			}
		}()

		if err = goose.Run(args[0], dbConn, "migrations", os.Args[1:]...); err != nil {
			log.Err(err).Send()
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
