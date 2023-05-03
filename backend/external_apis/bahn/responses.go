package bahn

import (
	"time"
)

type Place struct {
	ID string `json:"id"`
	// Full station name
	Name string `json:"name"`
	// If set label is usually a shorter name than Name
	Label string `json:"label"`
	// enum: "stop", "poi"
	Type string `json:"type"`
	// enum: "STATION", "SHOP", "PUBLIC", "STREET", "POI"
	Category string `json:"category"`
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
	Type      string     `json:"type"`
	Departure TimingInfo `json:"departure,omitempty"`
	Arrival   TimingInfo `json:"arrival,omitempty"`
	Duration  struct {
		Value int    `json:"value"`
		Unit  string `json:"unit"`
	} `json:"duration"`
	ID          string `json:"id"`
	Fingerprint string `json:"fingerprint"`
	Fare        struct {
		Warnings         []interface{} `json:"warnings"`
		TicketsAvailable bool          `json:"ticketsAvailable"`
	} `json:"fare"`
	Legs []struct {
		ID     string `json:"id"`
		Type   string `json:"type"`
		Origin struct {
			Coordinates struct {
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
			} `json:"coordinates"`
			Name              string        `json:"name"`
			ScheduledPlatform string        `json:"scheduledPlatform"`
			StationID         string        `json:"stationId"`
			Entrances         []interface{} `json:"entrances"`
		} `json:"origin"`
		Destination struct {
			Coordinates struct {
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
			} `json:"coordinates"`
			Name              string        `json:"name"`
			ScheduledPlatform string        `json:"scheduledPlatform"`
			StationID         string        `json:"stationId"`
			Entrances         []interface{} `json:"entrances"`
		} `json:"destination"`
		Departure struct {
			ScheduledTime time.Time `json:"scheduledTime"`
		} `json:"departure"`
		Arrival struct {
			ScheduledTime time.Time `json:"scheduledTime"`
		} `json:"arrival"`
		Duration struct {
			Value int    `json:"value"`
			Unit  string `json:"unit"`
		} `json:"duration"`
		Line struct {
			Name          string `json:"name"`
			HeadSign      string `json:"headSign"`
			Color         string `json:"color"`
			TextColor     string `json:"textColor"`
			ModalityColor string `json:"modalityColor"`
			Vehicle       string `json:"vehicle"`
			JourneyID     string `json:"journeyId"`
		} `json:"line"`
		Shape struct {
			Color string `json:"color"`
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"shape"`
		Distance struct {
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
		} `json:"distance"`
		IconUrls struct {
			IconURL string `json:"iconUrl"`
		} `json:"iconUrls"`
		AlternativeLegDepartures []interface{} `json:"alternativeLegDepartures"`
		GramsOfCO2Emitted        int           `json:"gramsOfCO2Emitted"`
		AlternativeLines         []interface{} `json:"alternativeLines"`
		IntermediateStops        []struct {
			Stop struct {
				Coordinates struct {
					Longitude float64 `json:"longitude"`
					Latitude  float64 `json:"latitude"`
				} `json:"coordinates"`
				Name              string        `json:"name"`
				ScheduledPlatform string        `json:"scheduledPlatform"`
				StationID         string        `json:"stationId"`
				Entrances         []interface{} `json:"entrances"`
			} `json:"stop"`
			ArrivalTime struct {
				ScheduledTime time.Time `json:"scheduledTime"`
			} `json:"arrivalTime"`
			DepartureTime struct {
				ScheduledTime time.Time `json:"scheduledTime"`
			} `json:"departureTime"`
			Cancelled bool `json:"cancelled"`
		} `json:"intermediateStops"`
		Messages        []interface{} `json:"messages"`
		Fingerprint     string        `json:"fingerprint"`
		TrainName       string        `json:"trainName"`
		TrainAttributes []string      `json:"trainAttributes"`
	} `json:"legs"`
	PublicTransportStart struct {
		PTDeparture struct {
			ScheduledTime time.Time `json:"scheduledTime"`
		} `json:"pTDeparture"`
		PTStationName string `json:"pTStationName"`
	} `json:"publicTransportStart"`
	ExtraCalls []struct {
		Href     string `json:"href"`
		Relation string `json:"relation"`
		Method   string `json:"method"`
	} `json:"extraCalls"`
	Messages          []interface{} `json:"messages"`
	GramsOfCO2Emitted int           `json:"gramsOfCO2Emitted"`
}

type ConnectionsResponse struct {
	Trips      []Trip `json:"trips"`
	ExtraCalls []struct {
		Href     string `json:"href"`
		Relation string `json:"relation"`
		Method   string `json:"method"`
	} `json:"extraCalls"`
	Warnings []interface{} `json:"warnings"`
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
