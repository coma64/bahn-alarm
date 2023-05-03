package watcher

import (
	"context"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/watcher/checking"
	"github.com/coma64/bahn-alarm-backend/watcher/queries"
	"github.com/rs/zerolog/log"
	"math/rand"
	"time"
)

func checkDueDepartures(ctx context.Context) error {
	departures, err := queries.SelectDueDepartures(ctx)
	if err != nil {
		return fmt.Errorf("error getting due departures: %w", err)
	}

	for _, departure := range departures {
		var newNextCheck time.Time
		if err = checking.CheckDeparture(ctx, &departure); err != nil {
			log.Err(err).Msg("Failed to check departure")
			newNextCheck = time.Now().UTC().Add(time.Hour)
		} else {
			newNextCheck = getNewNextCheck(&departure)
		}

		log.Debug().
			Int("departureId", departure.Id).
			Dur("nextCheckDiff", newNextCheck.Sub(time.Now().UTC())).
			Msg("Updating next check")

		if err = queries.UpdateDepartureNextCheck(ctx, departure.Id, newNextCheck); err != nil {
			return fmt.Errorf("error updating next check: %w", err)
		}
		departure.NextCheck = newNextCheck
	}

	return nil
}

func nowAddWithJitter(sleep time.Duration) time.Time {
	return time.Now().Add(sleep + time.Duration(rand.Int63n(int64(sleep*5/100))))
}

func getNewNextCheck(connection *queries.FatDeparture) time.Time {
	nextDeparture := connection.TimeUntilNextDeparture()
	if nextDeparture < time.Minute*2 {
		return nowAddWithJitter(nextDeparture)
	} else if nextDeparture < time.Minute*5 {
		return nowAddWithJitter(time.Minute)
	} else if nextDeparture < time.Minute*10 {
		return nowAddWithJitter(time.Minute * 2)
	} else if nextDeparture < time.Minute*20 {
		return nowAddWithJitter(time.Minute * 5)
	} else if nextDeparture < time.Hour {
		return nowAddWithJitter(time.Minute * 15)
	} else if nextDeparture < time.Hour*6 {
		return nowAddWithJitter(time.Minute * 30)
	} else if nextDeparture < time.Hour*12 {
		return nowAddWithJitter(time.Hour)
	} else {
		return nowAddWithJitter(time.Hour * 2)
	}
}
