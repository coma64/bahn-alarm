package checking

import (
	"context"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/metrics"
	"github.com/coma64/bahn-alarm-backend/notifications"
	"github.com/coma64/bahn-alarm-backend/notifications/web_push_notifier"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/coma64/bahn-alarm-backend/watcher/queries"
	"github.com/rs/zerolog/log"
)

func CheckDeparture(ctx context.Context, departure *queries.FatDeparture) (hasSentNotification, departureIsOnTime bool, err error) {
	log.Debug().
		Int("departureId", departure.Id).
		Time("departureTime", departure.Departure.Departure).
		Int("trackedById", departure.TrackedById).
		Msg("Checking departure")

	trip, err := fetchTrip(ctx, departure)
	if err != nil {
		return false, false, fmt.Errorf("error fetching trip: %w", err)
	}

	var newDepartureInfos *models.DepartureInfo
	newDepartureInfos, err = queries.CreateOrUpdateDepartureInfo(ctx, departure, trip)
	if err != nil {
		return false, false, fmt.Errorf("error upserting delay infos: %w", err)
	}

	oldStatus := departure.Status
	oldDelay := departure.DelayMinutes
	newStatus := newDepartureInfos.DepartureStatus()
	newDelay := newDepartureInfos.DelayMinutes()

	notificationThreshold := 0
	if err = db.Db.GetContext(ctx, &notificationThreshold, "select notificationthresholdminutes from users where id = $1", departure.TrackedById); err != nil {
		return false, false, fmt.Errorf("error getting notification threshold: %w", err)
	}

	if !shouldSendNotification(oldStatus, newStatus, oldDelay, newDelay, notificationThreshold) {
		log.Debug().Int("departureId", departure.Id).Msg("Not sending notification")
		return false, newStatus == server.OnTime, nil
	}
	log.Debug().Int("departureId", departure.Id).Msg("Sending notification")

	if err = sendAlarm(ctx, departure, oldStatus, newStatus, oldDelay, newDelay); err != nil {
		return false, newStatus == server.OnTime, err
	}

	return true, newStatus == server.OnTime, nil
}

func sendAlarm(ctx context.Context, departure *queries.FatDeparture, oldStatus server.TrackedDepartureStatus, newStatus server.TrackedDepartureStatus, oldDelay int, newDelay int) error {
	urgency, message := getDelayMessage(oldStatus, newStatus, oldDelay, newDelay)

	var alarm *models.Alarm
	var err error
	if alarm, err = models.InsertAlarm(ctx, departure.TrackedById, urgency, departure.Id, message); err != nil {
		return fmt.Errorf("error creating alarm: %w", err)
	}

	var notification *notifications.Notification
	if notification, err = alarm.ToPushNotification(ctx, db.Db); err != nil {
		return fmt.Errorf("error converting alarm %d for user %d to notification: %w", alarm.Id, alarm.ReceiverId, err)
	}

	metrics.AlarmsSent.WithLabelValues(string(alarm.Urgency)).Inc()

	go func() {
		if err = web_push_notifier.New(db.Db).SendNotification(ctx, *notification, alarm.ReceiverId); err != nil {
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

func shouldSendNotification(oldStatus, newStatus server.TrackedDepartureStatus, oldDelay, newDelay, notificationThreshold int) bool {
	if newStatus == server.Delayed && newDelay <= notificationThreshold {
		log.Debug().Int("delay", newDelay).Int("threshold", notificationThreshold).Str("status", string(newStatus)).Msg("Not sending notification because of notification threshold")
		return false
	}

	if oldStatus != newStatus && !isFirstCheckAndIsOnTime(oldStatus, newStatus) {
		return true
	}

	return hasDelayChanged(oldStatus, newStatus, oldDelay, newDelay)
}
