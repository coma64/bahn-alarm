package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	AlarmsSent = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bahn_alarm_alarms_sent_total",
		Help: "The total number of web push notifications sent",
	})
	BahnApiRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "bahn_alarm_bahn_api_requests_total",
		Help: "The amount of requests sent to the bahn API",
	})
	RequestDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "bahn_alarm_request_duration",
		Help: "API Request duration",
	})
)
