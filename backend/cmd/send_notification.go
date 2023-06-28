package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"net/http/httputil"
	"os"
)

var username, content string
var sendNotificationCmd = &cobra.Command{
	Use:   "send-notification",
	Short: "Send a notification to all web push subscriptions of a user",
	Run: func(cmd *cobra.Command, args []string) {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

		subs := []models.PushNotificationSub{}
		if err := db.Db.Select(
			&subs,
			"select p.* from pushNotificationSubs p join users u on u.id = p.ownerId where u.name = $1",
			username,
		); err != nil {
			log.Fatal().Err(err).Send()
		}

		if len(subs) == 0 {
			log.Fatal().Msg("No push subscriptions found")
		}

		for _, sub := range subs {
			subLog := log.With().Int("subId", sub.Id).Str("subName", sub.Name).Logger()

			webpushSub := &webpush.Subscription{}
			if err := json.Unmarshal(sub.RawSubscription, webpushSub); err != nil {
				log.Err(err).Int("subId", sub.Id).Str("subName", sub.Name).Msg("Unable to unmarshal push subscription. Skipping")
				continue
			}

			webpushResponse, err := webpush.SendNotification([]byte(content), webpushSub, &webpush.Options{
				Subscriber:      config.Conf.PushNotifications.Subject,
				TTL:             config.Conf.PushNotifications.Ttl,
				VAPIDPublicKey:  config.Conf.PushNotifications.VapidKeys.Public,
				VAPIDPrivateKey: config.Conf.PushNotifications.VapidKeys.Private,
			})
			if err != nil {
				subLog.Err(err).Str("content", content).Msg("Failed to send push notification. Skipping")
				continue
			} else {
				subLog.Info().Msg("Successfully sent push notification")
			}

			var dump []byte
			if dump, err = httputil.DumpResponse(webpushResponse, true); err != nil {
				subLog.Err(err).Msg("Failed to dump webpush response body")
			} else {
				fmt.Println("--- Response ---")
				fmt.Println(string(dump))
			}

			if err = webpushResponse.Body.Close(); err != nil {
				subLog.Err(err).Msg("Failed to close webpush response body")
			}
		}
	},
}

func init() {
	sendNotificationCmd.Flags().StringVarP(&username, "username", "u", "", "The user who will receive the notification")
	sendNotificationCmd.Flags().StringVarP(&content, "content", "c", "", "The web push subscription content. Must be a JSON object. Only title is required, see https://angular.io/api/service-worker/SwPush#usage-notes")
	if err := sendNotificationCmd.MarkFlagRequired("username"); err != nil {
		log.Err(err).Msg("Failed to make send-notification cmd flag username required")
	}
	if err := sendNotificationCmd.MarkFlagRequired("content"); err != nil {
		log.Err(err).Msg("Failed to make send-notification cmd flag content required")
	}

	rootCmd.AddCommand(sendNotificationCmd)
}
