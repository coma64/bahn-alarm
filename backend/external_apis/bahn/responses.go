package bahn

import (
	"time"
)

type Place struct {
	ID string `json:"id"`
	// Full station name
	Name string `json:"name"`
	// ID required for the connections endpoint
	StationID string `json:"stationId,omitempty"`
}

type PlacesResponse struct {
	Places     []Place `json:"places"`
	DataSource string  `json:"dataSource"`
}

type TimingInfo struct {
	ScheduledTime time.Time `json:"scheduledTime"`
	ActualTime    time.Time `json:"actualTime"`
}

func (t *TimingInfo) IsOnTime() bool {
	return t.ActualTime.IsZero() || t.ActualTime.Sub(t.ScheduledTime) == 0
}

type Trip struct {
	Departure TimingInfo `json:"departure,omitempty"`
	Arrival   TimingInfo `json:"arrival,omitempty"`
}

type ConnectionsResponse struct {
	Trips []Trip `json:"trips"`
}

func (r *ConnectionsResponse) ProcessResponse() {
	for i := range r.Trips {
		r.Trips[i].Departure.ScheduledTime = r.Trips[i].Departure.ScheduledTime.UTC()
		r.Trips[i].Departure.ActualTime = r.Trips[i].Departure.ActualTime.UTC()
		r.Trips[i].Arrival.ScheduledTime = r.Trips[i].Arrival.ScheduledTime.UTC()
		r.Trips[i].Arrival.ActualTime = r.Trips[i].Arrival.ActualTime.UTC()

		if r.Trips[i].Departure.ActualTime.IsZero() {
			r.Trips[i].Departure.ActualTime = r.Trips[i].Departure.ScheduledTime
		}
	}
}
