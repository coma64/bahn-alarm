package models

import (
	"github.com/coma64/bahn-alarm-backend/db"
	"time"
)

type Departure struct {
	db.IdModel
	ConnectionId int
	Departure    time.Time
	NextCheck    time.Time
}
