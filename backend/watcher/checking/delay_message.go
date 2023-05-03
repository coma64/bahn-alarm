package checking

import (
	"fmt"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/coma64/bahn-alarm-backend/watcher/queries"
)

func getDelayMessage(oldStatus, newStatus models.DepartureStatus, oldDelay, newDelay int) (server.Urgency, string) {
	if oldStatus == models.DepartureStatusDelayed && newStatus == models.DepartureStatusDelayed && oldDelay != newDelay {
		if oldDelay < newDelay {
			return server.Warn, fmt.Sprintf("Connection delay has increased to %dm", newDelay)
		} else {
			return server.Info, fmt.Sprintf("Connection delay has decreased to %dm", newDelay)
		}
	} else if (oldStatus == models.DepartureStatusDelayed || oldStatus == models.DepartureStatusCanceled) && newStatus == models.DepartureStatusOnTime {
		return server.Info, "Connection is no longer delayed"
	} else if oldStatus == models.DepartureStatusCanceled && newStatus == models.DepartureStatusDelayed {
		return server.Info, fmt.Sprintf("Connection is no longer canceled, now delayed by %dm", newDelay)
	}

	switch newStatus {
	case models.DepartureStatusCanceled:
		return server.Error, "Connection was cancelled"
	case models.DepartureStatusOnTime:
		return server.Error, "Connection is on time"
	case models.DepartureStatusDelayed:
		return server.Error, fmt.Sprintf("Connection is delayed by %dm", newDelay)
	}
	return server.Info, "Connection was not checked yet"
}

func createConnectionAlarmContent(departure *queries.DepartureModel, message string) *server.ConnectionAlarm {
	return &server.ConnectionAlarm{
		Connection: server.SimpleConnection{
			Departure: departure.Departure.Departure,
			FromName:  departure.FromStationName,
			ToName:    departure.ToStationName,
		},
		Message: message,
		Type:    server.ConnectionAlarmTypeConnectionAlarm,
	}
}
