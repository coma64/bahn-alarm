package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var password string
var isAdmin bool

var createUserCmd = &cobra.Command{
	Use:       "create-user",
	Short:     "Create a new User",
	Aliases:   []string{"c"},
	ValidArgs: []string{"username"},
	Args:      cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		if len(args[0]) == 0 {
			log.Fatal().Msg("Username cannot be an empty string")
		}

		if len(password) == 0 {
			if err := survey.AskOne(&survey.Password{Message: "Enter password:"}, &password); err != nil {
				log.Fatal().Err(err).Msg("Failed to read password")
			} else if len(password) == 0 {
				log.Fatal().Msg("The password cannot be an empty string")
			}
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to hash password")
		}

		var newUser models.User
		if err = db.Db.Get(
			&newUser,
			"insert into users (name, passwordHash, isAdmin) values ($1, $2, $3) returning *;",
			args[0],
			string(hashedPassword),
			isAdmin,
		); err != nil {
			log.Fatal().Err(err).Msg("Failed to create user")
		}

		log.Info().Str("name", newUser.Name).Int("id", newUser.Id).Msg("Created user")
	},
}

func init() {
	createUserCmd.Flags().StringVarP(&password, "password", "p", "", "")
	createUserCmd.Flags().BoolVarP(&isAdmin, "is-admin", "a", false, "Whether the user will have admin permissions")

	rootCmd.AddCommand(createUserCmd)
}
