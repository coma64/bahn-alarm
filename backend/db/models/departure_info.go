package models

import (
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/rs/zerolog/log"
	"time"
)

type DepartureInfo struct {
	DepartureId   int
	Day           time.Time
	ScheduledTime time.Time
	ActualTime    *time.Time
}

func (d *DepartureInfo) DepartureStatus() server.TrackedDepartureStatus {
	logEvent := log.Debug().
		Int("departureId", d.DepartureId).
		Time("scheduledTime", d.ScheduledTime)

	if d.ActualTime != nil {
		logEvent = logEvent.Time("actualTime", *d.ActualTime)
	}

	if d == nil {
		logEvent.Msg("DepartureInfo is nil. Setting status to NotChecked")
		return server.NotChecked
	} else if d.ActualTime == nil {
		logEvent.Msg("ActualTime is nil. Settings status to Canceled")
		return server.Canceled
	} else if d.ScheduledTime.Equal(*d.ActualTime) {
		logEvent.Msg("ScheduledTime is the same as ActualTime. Setting status to OnTime")
		return server.OnTime
	} else {
		logEvent.Msg("Else case. Setting status to Delayed")
		return server.Delayed
	}
}

func (d *DepartureInfo) DelayMinutes() int {
	if d == nil || d.ActualTime == nil {
		return 0
	}

	return int(d.ActualTime.Sub(d.ScheduledTime).Minutes())
}
