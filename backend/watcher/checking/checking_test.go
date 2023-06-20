package checking

import (
	"github.com/coma64/bahn-alarm-backend/server"
	"testing"
)

func TestShouldSendNotification(t *testing.T) {
	if shouldSendNotification(server.Canceled, server.Canceled, 0, 0) {
		t.Fatal("Would send notification although nothings changed")
	}

	if shouldSendNotification(server.NotChecked, server.OnTime, 0, 0) {
		t.Fatal("Would send notification although this is the first check")
	}

	if !shouldSendNotification(server.Delayed, server.Delayed, 3, 4) {
		t.Fatal("Not sending notification although delay has changed")
	}

	if shouldSendNotification(server.Delayed, server.Delayed, 4, 4) {
		t.Fatal("Sending notification although delay has not changed")
	}

	if !shouldSendNotification(server.OnTime, server.Delayed, 0, 0) {
		t.Fatal("Not sending notification although status changed")
	}
}
