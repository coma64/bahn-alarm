package models

import "time"

const (
	DepartureStatusNotChecked DepartureStatus = iota
	DepartureStatusDelayed
	DepartureStatusCanceled
	DepartureStatusOnTime
)

type DepartureStatus int

type DepartureInfo struct {
	DepartureId   int
	Day           time.Time
	ScheduledTime time.Time
	ActualTime    *time.Time
}

func (d *DepartureInfo) DepartureStatus() DepartureStatus {
	if d == nil {
		return DepartureStatusNotChecked
	} else if d.ActualTime == nil {
		return DepartureStatusCanceled
	} else if d.ScheduledTime.Equal(*d.ActualTime) {
		return DepartureStatusOnTime
	} else {
		return DepartureStatusDelayed
	}
}

func (d *DepartureInfo) DelayMinutes() int {
	if d == nil || d.ActualTime == nil {
		return 0
	}

	return int(d.ActualTime.Sub(d.ScheduledTime).Minutes())
}
