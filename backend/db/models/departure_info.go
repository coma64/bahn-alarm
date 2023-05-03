package models

import (
	"github.com/coma64/bahn-alarm-backend/server"
	"time"
)

type DepartureInfo struct {
	DepartureId   int
	Day           time.Time
	ScheduledTime time.Time
	ActualTime    *time.Time
}

func (d *DepartureInfo) DepartureStatus() server.TrackedDepartureStatus {
	if d == nil {
		return server.NotChecked
	} else if d.ActualTime == nil {
		return server.Canceled
	} else if d.ScheduledTime.Equal(*d.ActualTime) {
		return server.OnTime
	} else {
		return server.Delayed
	}
}

func (d *DepartureInfo) DelayMinutes() int {
	if d == nil || d.ActualTime == nil {
		return 0
	}

	return int(d.ActualTime.Sub(d.ScheduledTime).Minutes())
}
