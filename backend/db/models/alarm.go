package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/jmoiron/sqlx/types"
	"github.com/rs/zerolog/log"
	"time"
)

type AlarmUrgency string

type Alarm struct {
	db.IdModel
	ReceiverId int
	CreatedAt  time.Time
	Urgency    server.Urgency
	AlarmData  types.JSONText
}

func (a *Alarm) ToSchema() (*server.Alarm, error) {
	content := server.Alarm_Content{}
	if err := content.UnmarshalJSON(a.AlarmData); err != nil {
		return nil, err
	}

	return &server.Alarm{
		Content:   content,
		CreatedAt: a.CreatedAt,
		Id:        a.Id,
		Urgency:   a.Urgency,
	}, nil
}

func InsertAlarm(ctx context.Context, receiverId int, urgency server.Urgency, content *server.ConnectionAlarm) (*Alarm, error) {
	contentJson, err := json.Marshal(content)
	if err != nil {
		return nil, fmt.Errorf("error marshaling alarm content: %w", err)
	}

	alarm := &Alarm{}
	return alarm, db.Db.GetContext(
		ctx,
		alarm,
		"insert into alarms (receiverId, urgency, alarmData) values ($1, $2, $3) returning *",
		receiverId,
		urgency,
		contentJson,
	)
}

func (a *Alarm) SendPushNotification() error {
	var sub PushNotificationSub
	if err := db.Db.Get(
		&sub,
		"select p.* from pushNotificationSubs p join users u on u.id = p.ownerId where u.id = $1",
		a.ReceiverId,
	); err != nil {
		log.Debug().Int("receiverId", a.ReceiverId).Int("alarmId", a.Id).Msg("No push subscription found for alarm receiver")
		return nil
	}

	webpushSub := &webpush.Subscription{}
	if err := json.Unmarshal(sub.RawSubscription, &webpushSub); err != nil {
		return fmt.Errorf("error unmarshaling webpush subscription: %w", err)
	}

	go func() {
		_, err := webpush.SendNotification(a.AlarmData, webpushSub, &webpush.Options{
			Subscriber:      "coma64@outlook.com",
			TTL:             30,
			VAPIDPublicKey:  config.Conf.PushNotifications.VapidKeys.Public,
			VAPIDPrivateKey: config.Conf.PushNotifications.VapidKeys.Private,
		})
		if err != nil {
			log.Err(err).Int("receiverId", a.ReceiverId).Int("alarmId", a.Id).Msg("Failed to send push notification")
		}
	}()

	return nil
}
