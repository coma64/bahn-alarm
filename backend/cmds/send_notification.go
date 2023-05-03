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

	var sub models.PushNotificationSub
	if err := db.Db.Get(
		&sub,
		"select p.* from pushNotificationSubs p join users u on u.id = p.ownerId where u.name = $1",
		response.Name,
	); err != nil {
		log.Fatal().Err(err).Send()
	}

	webpushSub := &webpush.Subscription{}
	if err := json.Unmarshal(sub.RawSubscription, &webpushSub); err != nil {
		log.Fatal().Err(err).Send()
	}

	webpushResponse, err := webpush.SendNotification([]byte(response.Content), webpushSub, &webpush.Options{
		Subscriber:      "coma64@outlook.com",
		TTL:             30,
		VAPIDPublicKey:  config.Conf.PushNotifications.VapidKeys.Public,
		VAPIDPrivateKey: config.Conf.PushNotifications.VapidKeys.Private,
	})
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	if dump, err := httputil.DumpResponse(webpushResponse, true); err != nil {
		log.Fatal().Err(err).Send()
	} else {
		fmt.Printf("%q", dump)
	}

	if err = webpushResponse.Body.Close(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
