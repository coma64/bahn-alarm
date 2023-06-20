package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http/httputil"
	"os"
)

var questions = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Enter username:"},
		Validate: survey.Required,
	},
	{
		Name:     "content",
		Prompt:   &survey.Input{Message: "Enter message content:"},
		Validate: survey.Required,
	},
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	response := struct {
		Name    string
		Content string
	}{}
	if err := survey.Ask(questions, &response); err != nil {
		log.Fatal().Err(err).Send()
	}

	subs := []models.PushNotificationSub{}
	if err := db.Db.Select(
		&subs,
		"select p.* from pushNotificationSubs p join users u on u.id = p.ownerId where u.name = $1",
		response.Name,
	); err != nil {
		log.Fatal().Err(err).Send()
	}

	for _, sub := range subs {
		subLog := log.With().Int("subId", sub.Id).Str("subName", sub.Name).Logger()

		webpushSub := &webpush.Subscription{}
		if err := json.Unmarshal(sub.RawSubscription, webpushSub); err != nil {
			log.Err(err).Int("subId", sub.Id).Str("subName", sub.Name).Msg("Unable to unmarshal push subscription. Skipping")
			continue
		}

		webpushResponse, err := webpush.SendNotification([]byte(response.Content), webpushSub, &webpush.Options{
			Subscriber:      config.Conf.PushNotifications.Subject,
			TTL:             config.Conf.PushNotifications.Ttl,
			VAPIDPublicKey:  config.Conf.PushNotifications.VapidKeys.Public,
			VAPIDPrivateKey: config.Conf.PushNotifications.VapidKeys.Private,
		})
		if err != nil {
			subLog.Err(err).Str("content", response.Content).Msg("Failed to send push notification. Skipping")
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
}
