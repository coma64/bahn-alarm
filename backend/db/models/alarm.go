package models

import (
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/jmoiron/sqlx/types"
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
