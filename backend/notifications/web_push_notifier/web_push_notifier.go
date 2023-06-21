package web_push_notifier

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/coma64/bahn-alarm-backend/config"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/notifications"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"net/http"
	"net/http/httputil"
)

type WebPushNotifier struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) WebPushNotifier {
	return WebPushNotifier{db: db}
}

func (w WebPushNotifier) SendNotification(ctx context.Context, notification notifications.Notification, userId int) error {
	webPushSubscriptions, err := models.SelectWebPushSubscriptionForUser(ctx, w.db, userId)
	if err != nil {
		return fmt.Errorf("error getting web push subscriptions: %w", err)
	}

	var notificationJson []byte
	notificationJson, err = json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("error marshaling notification: %w", err)
	}

	for _, sub := range webPushSubscriptions {
		var response *http.Response
		if response, err = webpush.SendNotificationWithContext(
			ctx,
			notificationJson,
			&sub.WebPushSubscription,
			&webpush.Options{
				Subscriber:      config.Conf.PushNotifications.Subject,
				TTL:             config.Conf.PushNotifications.Ttl,
				VAPIDPublicKey:  config.Conf.PushNotifications.VapidKeys.Public,
				VAPIDPrivateKey: config.Conf.PushNotifications.VapidKeys.Private,
			},
		); err != nil {
			return fmt.Errorf("error sending push notification to web push subscription %d from user %d: %w", sub.Id, userId, err)
		}

		if err = w.handleResponse(ctx, response, sub); err != nil {
			return err
		}

		if err = response.Body.Close(); err != nil {
			return fmt.Errorf("error closing response body from web push subscription %d from user %d: %w", sub.Id, userId, err)
		}
	}

	return nil
}

func (w WebPushNotifier) handleResponse(ctx context.Context, response *http.Response, sub models.UnmarshaledPushNotificationSub) error {
	subLog := log.With().Int("status", response.StatusCode).Int("pushNotificationSubId", sub.Id).Int("userId", sub.OwnerId).Logger()

	if response.StatusCode == 410 {
		subLog.Debug().Msg("Tried to send web push notification to old subscription")

		if err := sub.PushNotificationSub.Delete(ctx, w.db); err != nil {
			return fmt.Errorf("error removing old push notification subscription: %w", err)
		}
	} else if response.StatusCode == 201 {
		subLog.Debug().Msg("Sent web push notification")
	} else {
		subLog.Error().Msg("Got unexpected response when sending web push notification. See log for full response")

		if dump, err := httputil.DumpResponse(response, true); err != nil {
			subLog.Err(err).Msg("Failed to dump web push notification response body")
		} else {
			fmt.Println("--- Response ---")
			fmt.Println(string(dump))
		}

		return fmt.Errorf("got unexpected response status code %d when sending web push notification to subscription %d for user %d", response.StatusCode, sub.Id, sub.OwnerId)
	}

	return nil
}
