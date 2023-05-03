package main

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"os"
)

var questions = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Enter username:"},
		Validate: survey.Required,
	},
	{
		Name:   "password",
		Prompt: &survey.Password{Message: "Enter password:"},
	},
	{
		Name:   "isAdmin",
		Prompt: &survey.Confirm{Message: "Should the user be an admin?", Default: false},
	},
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	response := struct {
		Name     string
		Password string
		IsAdmin  bool
	}{}
	if err := survey.Ask(questions, &response); err != nil {
		log.Fatal().Err(err).Send()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(response.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to hash password")
	}

	var newUser models.User
	if err = db.Db.Get(
		&newUser,
		"insert into users (name, passwordHash, isAdmin) values ($1, $2, $3) returning *;",
		response.Name,
		string(hashedPassword),
		response.IsAdmin,
	); err != nil {
		log.Fatal().Err(err).Msg("Failed to create user")
	}

	log.Info().Str("name", newUser.Name).Int("id", newUser.Id).Msg("Created user")
}
