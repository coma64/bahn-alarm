package checking

import (
	"context"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/watcher/queries"
	"github.com/rs/zerolog/log"
)

func CheckDeparture(ctx context.Context, departure *queries.DepartureModel) error {
	log.Debug().
		Int("departureId", departure.Id).
		Time("departureTime", departure.Departure.Departure).
		Int("trackedById", departure.TrackedById).
		Msg("Checking departure")

	trip, err := fetchTrip(ctx, departure)
	if err != nil {
		return fmt.Errorf("error fetching trip: %w", err)
	}

	var oldDepartureInfos *models.DepartureInfo
	oldDepartureInfos, err = queries.GetDepartureInfos(ctx, departure.Id)
	if err != nil {
		return fmt.Errorf("error getting existing departure infos")
	}

	var newDepartureInfos *models.DepartureInfo
	newDepartureInfos, err = queries.CreateOrUpdateDepartureInfo(ctx, departure, trip)
	if err != nil {
		return fmt.Errorf("error upserting delay infos: %w", err)
	}

	oldStatus := oldDepartureInfos.DepartureStatus()
	oldDelay := oldDepartureInfos.DelayMinutes()
	newStatus := newDepartureInfos.DepartureStatus()
	newDelay := newDepartureInfos.DelayMinutes()
	if !shouldSendNotification(oldStatus, newStatus, oldDelay, newDelay) {
		log.Debug().Int("departureId", departure.Id).Msg("Not sending notification")
		return nil
	}
	log.Debug().Int("departureId", departure.Id).Msg("Sending notification")

	urgency, message := getDelayMessage(oldStatus, newStatus, oldDelay, newDelay)
	alarmContent := createConnectionAlarmContent(departure, message)

	var alarm *models.Alarm
	if alarm, err = models.InsertAlarm(ctx, departure.TrackedById, urgency, alarmContent); err != nil {
		return fmt.Errorf("error creating alarm: %w", err)
	}

	if err = alarm.SendPushNotification(); err != nil {
		return fmt.Errorf("error sending alarm as push notification: %w", err)
	}

	return nil
}

func hasDelayChanged(oldStatus, newStatus models.DepartureStatus, oldDelay, newDelay int) bool {
	return oldStatus == models.DepartureStatusDelayed && newStatus == models.DepartureStatusDelayed && oldDelay != newDelay
}

func isFirstCheckAndIsOnTime(oldStatus, newStatus models.DepartureStatus) bool {
	return oldStatus == models.DepartureStatusNotChecked && newStatus == models.DepartureStatusOnTime
}

func shouldSendNotification(oldStatus, newStatus models.DepartureStatus, oldDelay, newDelay int) bool {
	return isFirstCheckAndIsOnTime(oldStatus, newStatus) || (oldStatus == newStatus && !hasDelayChanged(oldStatus, newStatus, oldDelay, newDelay))
}
