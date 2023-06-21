package models

import (
	"context"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/notifications"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/jmoiron/sqlx"
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

var germanTimezone *time.Location

func init() {
	loc, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(fmt.Errorf("unable to get german timezone: %w", err))
	}

	germanTimezone = loc
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

func (a *Alarm) getNotificationDeparture(ctx context.Context, db *sqlx.DB) (string, error) {
	var departure time.Time
	if row := db.QueryRowxContext(ctx, "select departure from departures where id = $1", a.DepartureId); row.Err() != nil {
		return "", fmt.Errorf("error getting departure time: %w", row.Err())
	} else if err := row.Scan(&departure); err != nil {
		return "", fmt.Errorf("error scanning departure time: %w", err)
	}

	// TODO: convert to users actual timezone
	departure = departure.In(germanTimezone)

	return departure.Format("15:04"), nil
}

func (a *Alarm) ToPushNotification(ctx context.Context, db *sqlx.DB) (*notifications.Notification, error) {
	stations := struct {
		FromStationName string
		ToStationName   string
	}{}
	if err := db.GetContext(
		ctx,
		&stations,
		"select fromStationName, toStationname from fatDepartures where id = $1",
		a.DepartureId,
	); err != nil {
		return nil, fmt.Errorf("error getting station names: %w", err)
	}

	departure, err := a.getNotificationDeparture(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("error getting alarm departure for alarm %d on departure %d: %w", a.Id, a.DepartureId, err)
	}

	return &notifications.Notification{
		Title: departure + " " + stations.FromStationName + " -> " + stations.ToStationName,
		Body:  a.Message,
		Data: &notifications.NotificationData{
			OnActionClick: map[string]notifications.ActionClickOperation{
				"default": {Operation: "openWindow", Url: "/alarms"},
			},
		},
	}, nil
}
