package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

var (
	Prefix     = "bahn_alarm"
	AlarmsSent = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: Prefix + "_alarms_sent_total",
			Help: "The total number of web push notifications sent",
		},
		[]string{"urgency"},
	)
	PushApiRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: Prefix + "_push_api_request_duration_seconds",
			Help: "The time it took the web push API, of the respective vendor, to answer",
		},
		[]string{"status_code", "host"},
	)
	BahnApiRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: Prefix + "_bahn_api_request_duration_seconds",
			Help: "The time it took the Bahn API to respond",
		},
		[]string{"status_code", "path"},
	)
	WatcherCheckDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: Prefix + "_watcher_check_duration_seconds",
			Help: "The time the watcher took to check whether a particular departure was delayed",
		},
		[]string{"has_sent_notification", "departure_is_not_on_time", "fromStationName", "toStationName", "departureTime", "check_was_successful"},
	)
)

func AccurateSecondsSince(start time.Time) float64 {
	return float64(time.Now().Sub(start)) / float64(time.Second)
}
