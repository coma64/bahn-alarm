package checking

import (
	"context"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/external_apis/bahn"
	"github.com/coma64/bahn-alarm-backend/watcher/queries"
	"time"
)

var TripNotFound = fmt.Errorf("trip not found")

func fetchTrip(ctx context.Context, departure *queries.FatDeparture) (*bahn.Trip, error) {
	response, err := bahn.FetchConnections(
		ctx,
		time.Now().UTC().Add(
			departure.TimeUntilNextDeparture()-time.Duration(departure.DepartureMarginMinutes)*time.Minute,
		),
		departure.FromStationId,
		departure.ToStationId,
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching connections: %w", err)
	}

	response.ProcessResponse()

	var trip *bahn.Trip
	trip, err = getMatchingTrip(response.Trips, departure)
	if err != nil {
		return nil, err
	}

	return trip, nil
}

func getMatchingTrip(trips []bahn.Trip, departure *queries.FatDeparture) (*bahn.Trip, error) {
	for _, trip := range trips {
		if isWithinMargin(trip.Departure.ScheduledTime, departure.Departure.Departure, departure.DepartureMarginMinutes) {
			return &trip, nil
		}
	}
	return nil, TripNotFound
}

func isWithinMargin(tripDeparture, expectedDeparture time.Time, departureMarginMinutes int) bool {
	tripTime := queries.TimeOnly(tripDeparture)
	minimumDeparture := expectedDeparture.Add(time.Duration(-departureMarginMinutes) * time.Minute)
	maximumDeparture := expectedDeparture.Add(time.Duration(departureMarginMinutes) * time.Minute)
	return tripTime.After(minimumDeparture) && tripTime.Before(maximumDeparture)
}
