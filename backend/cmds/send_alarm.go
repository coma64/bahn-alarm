package main

import (
	"context"
	"github.com/AlecAivazis/survey/v2"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	var alarmIdStr string
	if err := survey.AskOne(&survey.Input{Message: "Enter alarm id:"}, &alarmIdStr); err != nil {
		log.Fatal().Err(err).Send()
	}

	alarmId, err := strconv.Atoi(alarmIdStr)
	if err != nil {
		log.Fatal().Msg("Passed in alarm id is not a number")
	}

	var alarm models.Alarm
	if err = db.Db.Get(
		&alarm,
		"select * from alarms where id = $1",
		alarmId,
	); err != nil {
		log.Fatal().Err(err).Msg("Cannot find alarm")
	}

	if err = alarm.SendPushNotification(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("Failed to send push notification")
	}

	log.Info().Msg("Sent push notification")
}
