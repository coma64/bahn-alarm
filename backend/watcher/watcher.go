package watcher

import (
	"context"
	"fmt"
	"github.com/coma64/bahn-alarm-backend/watcher/queries"
	"github.com/rs/zerolog/log"
	"time"
)

func WatchBahnApi(ctx context.Context) {
	// TODO: signal for new connections
	consecutiveErrorCount := 0
	var lastErr error
	for {
		if lastErr != nil {
			consecutiveErrorCount += 1
			errorBackoffSleep(consecutiveErrorCount, lastErr)
		} else {
			consecutiveErrorCount = 0
		}

		if lastErr = sleepUntilNextCheck(ctx); lastErr != nil {
			continue
		}

		if lastErr = checkDueDepartures(ctx); lastErr != nil {
			continue
		}
	}
}

func sleepUntilNextCheck(ctx context.Context) error {
	if nextCheck, err := queries.GetNextCheck(ctx); err != nil {
		return fmt.Errorf("error getting next check timestamp: %w", err)
	} else {
		duration := nextCheck.Sub(time.Now().UTC())
		log.Debug().Time("nextCheck", *nextCheck).Dur("duration", duration).Msg("Sleeping until next check")
		time.Sleep(duration)
	}

	return nil
}

func errorBackoffSleep(consecutiveErrorCount int, lastErr error) {
	backoffMinutes := time.Duration(consecutiveErrorCount) * 5
	if backoffMinutes > 30 {
		backoffMinutes = 30
	}

	backoffDuration := backoffMinutes * time.Minute
	log.Err(lastErr).Dur("backoffDuration", backoffDuration).Msg("Failed to check connections. Sleeping")

	time.Sleep(backoffDuration)
}
