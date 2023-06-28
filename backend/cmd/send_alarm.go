package cmd

import (
	"context"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/notifications"
	"github.com/coma64/bahn-alarm-backend/notifications/web_push_notifier"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"os"
)

var alarmId int

var sendAlarmCmd = &cobra.Command{
	Use:   "send-alarm",
	Short: "Send a already created alarm by Id",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		var alarm models.Alarm
		if err := db.Db.Get(
			&alarm,
			"select * from alarms where id = $1",
			alarmId,
		); err != nil {
			log.Fatal().Err(err).Msg("Cannot find alarm")
		}

		var err error
		var notification *notifications.Notification
		if notification, err = alarm.ToPushNotification(context.Background(), db.Db); err != nil {
			log.Fatal().Err(err).Msg("Failed to convert alarm to notification")
		}

		if err := web_push_notifier.New(db.Db).SendNotification(context.Background(), *notification, alarm.ReceiverId); err != nil {
			log.Fatal().Err(err).Msg("Failed to send push notification")
		}

		log.Info().Msg("Finished sending push notifications")
	},
}

func init() {
	sendAlarmCmd.Flags().IntVarP(&alarmId, "alarm-id", "i", 0, "")
	if err := sendAlarmCmd.MarkFlagRequired("alarm-id"); err != nil {
		log.Err(err).Msg("Failed to make alarm cmd flag alarm-id required")
	}
	rootCmd.AddCommand(sendAlarmCmd)
}
