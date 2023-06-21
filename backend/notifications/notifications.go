package notifications

import (
	"context"
)

type Notifier interface {
	SendNotification(ctx context.Context, notification Notification, userId int) error
}

// Notification https://developer.mozilla.org/en-US/docs/Web/API/Notification/Notification
type Notification struct {
	Title   string            `json:"title"`
	Body    string            `json:"body"`
	Actions []Action          `json:"actions,omitempty"`
	Data    *NotificationData `json:"data,omitempty"`
}

type NotificationData struct {
	OnActionClick map[string]ActionClickOperation `json:"onActionClick"`
}

type ActionClickOperation struct {
	Operation string `json:"operation"`
	Url       string `json:"url"`
}

type Action struct {
	Action string
	Title  string
}
