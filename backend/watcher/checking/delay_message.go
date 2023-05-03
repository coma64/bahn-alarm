package checking

import (
	"fmt"
	"github.com/coma64/bahn-alarm-backend/server"
)

func getDelayMessage(oldStatus, newStatus server.TrackedDepartureStatus, oldDelay, newDelay int) (server.Urgency, string) {
	if oldStatus == server.Delayed && newStatus == server.Delayed && oldDelay != newDelay {
		if oldDelay < newDelay {
			return server.Warn, fmt.Sprintf("Connection delay has increased to %dm", newDelay)
		} else {
			return server.Info, fmt.Sprintf("Connection delay has decreased to %dm", newDelay)
		}
	} else if oldStatus == server.Delayed && newStatus == server.OnTime {
		return server.Info, "Connection is no longer delayed"
	} else if oldStatus == server.Canceled {
		if newStatus == server.Delayed {
			return server.Warn, fmt.Sprintf("Connection is no longer canceled, now delayed by %dm", newDelay)
		} else {
			return server.Warn, fmt.Sprintf("Connection is no longer canceled")
		}
	}

	switch newStatus {
	case server.Canceled:
		return server.Error, "Connection was cancelled"
	case server.OnTime:
		return server.Info, "Connection is on time"
	case server.Delayed:
		urgency := server.Warn
		if newDelay > 5 {
			urgency = server.Error
		}

		return urgency, fmt.Sprintf("Connection is delayed by %dm", newDelay)
	}
	return server.Info, "Connection was not checked yet"
}
