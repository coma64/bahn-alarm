package models

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	"time"
)

type PushNotificationSub struct {
	db.IdModel
	OwnerId         int
	CreatedAt       time.Time
	RawSubscription types.JSONText
	Name            string
	IsEnabled       bool
}

func (s PushNotificationSub) ToSchema() (*server.PushNotificationSubscription, error) {
	var rawSub server.RawSubscription
	if err := s.RawSubscription.Unmarshal(&rawSub); err != nil {
		return nil, err
	}

	return &server.PushNotificationSubscription{
		CreatedAt:    &s.CreatedAt,
		Id:           &s.Id,
		IsEnabled:    s.IsEnabled,
		Name:         s.Name,
		Subscription: rawSub,
	}, nil
}

type UnmarshaledPushNotificationSub struct {
	PushNotificationSub
	WebPushSubscription webpush.Subscription
}

func SelectWebPushSubscriptionForUser(ctx context.Context, db *sqlx.DB, userId int) ([]UnmarshaledPushNotificationSub, error) {
	var subscriptions []PushNotificationSub
	if err := db.SelectContext(
		ctx,
		&subscriptions,
		"select p.* from pushNotificationSubs p join users u on u.id = p.ownerId where u.id = $1",
		userId,
	); err != nil {
		return nil, fmt.Errorf("error selecting push notification subscriptions for user %d: %w", userId, err)
	}

	var webPushSubscriptions []UnmarshaledPushNotificationSub
	for _, sub := range subscriptions {
		webPushSub := UnmarshaledPushNotificationSub{
			PushNotificationSub: sub,
			WebPushSubscription: webpush.Subscription{},
		}

		if err := json.Unmarshal(sub.RawSubscription, &webPushSub.WebPushSubscription); err != nil {
			return webPushSubscriptions, fmt.Errorf("error unmarshaling webpush subscription %d from user %d: %w", sub.Id, userId, err)
		}

		webPushSubscriptions = append(webPushSubscriptions, webPushSub)
	}

	return webPushSubscriptions, nil
}

func (p PushNotificationSub) Delete(ctx context.Context, db *sqlx.DB) error {
	if _, err := db.ExecContext(ctx, "delete from pushNotificationSubs where id = $1", p.Id); err != nil {
		return fmt.Errorf("error deleting push notification sub %d from user %d: %w", p.Id, p.OwnerId, err)
	}

	return nil
}
