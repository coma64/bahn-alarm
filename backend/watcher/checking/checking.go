package checking

import (
	"context"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/coma64/bahn-alarm-backend/watcher/queries"
	"github.com/rs/zerolog/log"
)

func CheckDeparture(ctx context.Context, departure *queries.FatDeparture) error {
	log.Debug().
		Int("departureId", departure.Id).
		Time("departureTime", departure.Departure.Departure).
		Int("trackedById", departure.TrackedById).
		Msg("Checking departure")

	trip, err := fetchTrip(ctx, departure)
	if err != nil {
		return fmt.Errorf("error fetching trip: %w", err)
	}

	var newDepartureInfos *models.DepartureInfo
	newDepartureInfos, err = queries.CreateOrUpdateDepartureInfo(ctx, departure, trip)
	if err != nil {
		return fmt.Errorf("error upserting delay infos: %w", err)
	}

	oldStatus := departure.Status
	oldDelay := departure.DelayMinutes
	newStatus := newDepartureInfos.DepartureStatus()
	newDelay := newDepartureInfos.DelayMinutes()
	if !shouldSendNotification(oldStatus, newStatus, oldDelay, newDelay) {
		log.Debug().Int("departureId", departure.Id).Msg("Not sending notification")
		return nil
	}
	log.Debug().Int("departureId", departure.Id).Msg("Sending notification")

	urgency, message := getDelayMessage(oldStatus, newStatus, oldDelay, newDelay)

	var alarm *models.Alarm
	if alarm, err = models.InsertAlarm(ctx, departure.TrackedById, urgency, departure.Id, message); err != nil {
		return fmt.Errorf("error creating alarm: %w", err)
	}

	go func() {
		if err = alarm.SendPushNotification(ctx); err != nil {
			log.Err(err).
				Int("receiverId", alarm.ReceiverId).
				Int("alarmId", alarm.Id).
				Msg("Failed to send push notification")
		}
	}()

	return nil
}

func hasDelayChanged(oldStatus, newStatus server.TrackedDepartureStatus, oldDelay, newDelay int) bool {
	return oldStatus == server.Delayed && newStatus == server.Delayed && oldDelay != newDelay
}

func isFirstCheckAndIsOnTime(oldStatus, newStatus server.TrackedDepartureStatus) bool {
	return oldStatus == server.NotChecked && newStatus == server.OnTime
}

func shouldSendNotification(oldStatus, newStatus server.TrackedDepartureStatus, oldDelay, newDelay int) bool {
	return !isFirstCheckAndIsOnTime(oldStatus, newStatus) || oldStatus != newStatus || !hasDelayChanged(oldStatus, newStatus, oldDelay, newDelay)
}
