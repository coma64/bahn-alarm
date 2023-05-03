package handlers

import (
	"context"
	"github.com/coma64/bahn-alarm-backend/server"
)

func (b *BahnAlarmApi) GetNotificationsPushSubscriptions(ctx context.Context, request server.GetNotificationsPushSubscriptionsRequestObject) (server.GetNotificationsPushSubscriptionsResponseObject, error) {
	return server.GetNotificationsPushSubscriptions200JSONResponse{}, nil
}

func (b *BahnAlarmApi) PostNotificationsPushSubscriptions(ctx context.Context, request server.PostNotificationsPushSubscriptionsRequestObject) (server.PostNotificationsPushSubscriptionsResponseObject, error) {
	return server.PostNotificationsPushSubscriptions201JSONResponse{}, nil
}

func (b *BahnAlarmApi) DeleteNotificationsPushSubscriptionsId(ctx context.Context, request server.DeleteNotificationsPushSubscriptionsIdRequestObject) (server.DeleteNotificationsPushSubscriptionsIdResponseObject, error) {
	return server.DeleteNotificationsPushSubscriptionsId204Response{}, nil
}

func (b *BahnAlarmApi) PatchNotificationsPushSubscriptionsId(ctx context.Context, request server.PatchNotificationsPushSubscriptionsIdRequestObject) (server.PatchNotificationsPushSubscriptionsIdResponseObject, error) {
	return server.PatchNotificationsPushSubscriptionsId204Response{}, nil
}
