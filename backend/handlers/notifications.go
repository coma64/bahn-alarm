package handlers

import (
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
)

func (b *BahnAlarmApi) GetNotificationsPushSubscriptions(ctx echo.Context, params server.GetNotificationsPushSubscriptionsParams) error {
	return nil
}

func (b *BahnAlarmApi) PostNotificationsPushSubscriptions(ctx echo.Context) error {
	return nil
}

func (b *BahnAlarmApi) DeleteNotificationsPushSubscriptionsId(ctx echo.Context, id int) error {
	return nil
}

func (b *BahnAlarmApi) PatchNotificationsPushSubscriptionsId(ctx echo.Context, id int) error {
	return nil
}
