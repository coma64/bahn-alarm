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
	BahnApiRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "bahn_alarm_bahn_api_requests_total",
			Help: "The amount of requests sent to the bahn API",
		},
		[]string{"path", "status_code"},
	)
)
