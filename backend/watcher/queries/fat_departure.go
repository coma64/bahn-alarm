package queries

import (
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/server"
	"time"
)

type FatDeparture struct {
	models.Departure
	FromStationId          string
	FromStationName        string
	ToStationId            string
	ToStationName          string
	DepartureMarginMinutes int
	TrackedById            int
	DelayMinutes           int
	Status                 server.TrackedDepartureStatus
}

func TimeOnly(t time.Time) time.Time {
	year, month, day := t.Date()
	return t.AddDate(-year, -int(month)+1, -day+1)
}

func (d *FatDeparture) HasDepartedToday() bool {
	nowTime := TimeOnly(time.Now().UTC())
	return d.Departure.Departure.Before(nowTime)
}

func (d *FatDeparture) TimeUntilNextDeparture() time.Duration {
	nowTime := TimeOnly(time.Now().UTC())
	diff := d.Departure.Departure.Sub(nowTime)
	if d.HasDepartedToday() {
		diff += time.Hour * 24
	}
	return diff
}
