package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/rs/zerolog/log"
	"time"
)

type AlarmUrgency string

type Alarm struct {
	db.IdModel
	ReceiverId  int
	CreatedAt   time.Time
	Urgency     server.Urgency
	DepartureId int
	Message     string
}

// https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification
type notificationOptions struct {
	Title string      `json:"title"`
	Body  string      `json:"body"`
	Data  interface{} `json:"data"`
}

type pushNotification struct {
	Notification notificationOptions `json:"notification"`
}

func (a *Alarm) ToSchema(connection *server.SimpleConnection) (*server.Alarm, error) {
	return &server.Alarm{
		Connection: *connection,
		CreatedAt:  a.CreatedAt,
		Id:         a.Id,
		Message:    a.Message,
		Urgency:    a.Urgency,
	}, nil
}

func InsertAlarm(ctx context.Context, receiverId int, urgency server.Urgency, departureId int, message string) (*Alarm, error) {
	alarm := &Alarm{}
	return alarm, db.Db.GetContext(
		ctx,
		alarm,
		"insert into alarms (receiverId, urgency, departureId, message) values ($1, $2, $3, $4) returning *",
		receiverId,
		urgency,
		departureId,
		message,
	)
}

func (a *Alarm) SendPushNotification(ctx context.Context) error {
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

	notification, err := a.toPushNotification(ctx)
	if err != nil {
		return fmt.Errorf("error creating push notification: %w", err)
	}

	var notificationJson []byte
	notificationJson, err = json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("error marshaling push notification: %w", err)
	}

	_, err = webpush.SendNotificationWithContext(ctx, notificationJson, webpushSub, &webpush.Options{
		Subscriber:      config.Conf.PushNotifications.Subject,
		TTL:             config.Conf.PushNotifications.Ttl,
		VAPIDPublicKey:  config.Conf.PushNotifications.VapidKeys.Public,
		VAPIDPrivateKey: config.Conf.PushNotifications.VapidKeys.Private,
	})

	return nil
}

func (a *Alarm) toPushNotification(ctx context.Context) (*pushNotification, error) {
	stations := struct {
		FromStationName string
		ToStationName   string
	}{}
	if err := db.Db.GetContext(
		ctx,
		&stations,
		"select fromStationName, toStationname from fatDepartures where id = $1",
		a.DepartureId,
	); err != nil {
		return nil, fmt.Errorf("error getting station names: %w", err)
	}

	return &pushNotification{
		Notification: notificationOptions{
			Title: stations.FromStationName + " -> " + stations.ToStationName,
			Body:  a.Message,
		},
	}, nil
}
