package models

import (
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/server"
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
