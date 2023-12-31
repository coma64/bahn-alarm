package cmd

import (
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

// See goose.run for available sub commands
var migrateCmd = &cobra.Command{
	Use:                "migrate",
	Short:              "Runs goose with the passed in arguments",
	Long:               "Runs the goose command with a predefined connection. See goose.run for available sub commands",
	Args:               cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	Aliases:            []string{"m"},

	Run: func(cmd *cobra.Command, args []string) {
		goose.SetBaseFS(migrations.EmbeddedMigrations)

		dbConn, err := goose.OpenDBWithDriver("postgres", db.CreateConnectionStr())
		if err != nil {
			log.Err(err).Msg("failed to open db")
		}

		defer func() {
			if err = dbConn.Close(); err != nil {
				log.Err(err).Msg("failed to close db")
			}
		}()

		if err = os.Chdir("migrations"); err != nil {
			log.Fatal().Err(err).Msg("Failed to cd into migration directory")
		}

		if err = goose.Run(args[0], dbConn, "migrations", args[1:]...); err != nil {
			log.Fatal().Err(err).Send()
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
