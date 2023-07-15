package queries

import (
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/coma64/bahn-alarm-backend/time_conversion"
	"time"
)

type StationInfo struct {
	FromStationId   string
	FromStationName string
	ToStationId     string
	ToStationName   string
}

type FatDeparture struct {
	models.Departure
	StationInfo
	DepartureMarginMinutes int
	TrackedById            int
	DelayMinutes           int
	Status                 server.TrackedDepartureStatus
}

func (d *FatDeparture) HasDepartedToday() bool {
	nowTime := time_conversion.TimeOnly(time.Now().UTC())
	return d.Departure.Departure.Before(nowTime)
}

func (d *FatDeparture) TimeUntilNextDeparture() time.Duration {
	nowTime := time_conversion.TimeOnly(time.Now().UTC())
	diff := d.Departure.Departure.Sub(nowTime)
	if d.HasDepartedToday() {
		diff += time.Hour * 24
	}
	return diff
}

func (f *FatDeparture) ToTrackedDeparture() server.TrackedDeparture {
	return server.TrackedDeparture{
		Delay:     f.DelayMinutes,
		Departure: f.Departure.Departure,
		Status:    f.Status,
	}
}
