package main

import (
	"context"
	"github.com/AlecAivazis/survey/v2"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/notifications"
	"github.com/coma64/bahn-alarm-backend/notifications/web_push_notifier"
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
		log.Fatal().Err(err).Msg("Passed in alarm id is not a number")
	}

	var alarm models.Alarm
	if err = db.Db.Get(
		&alarm,
		"select * from alarms where id = $1",
		alarmId,
	); err != nil {
		log.Fatal().Err(err).Msg("Cannot find alarm")
	}

	var notification *notifications.Notification
	if notification, err = alarm.ToPushNotification(context.Background(), db.Db); err != nil {
		log.Fatal().Err(err).Msg("Failed to convert alarm to notification")
	}

	if err = web_push_notifier.New(db.Db).SendNotification(context.Background(), *notification, alarm.ReceiverId); err != nil {
		log.Fatal().Err(err).Msg("Failed to send push notification")
	}

	log.Info().Msg("Finished sending push notifications")
}
