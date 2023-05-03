package handlers

import (
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
)

func (b *BahnAlarmApi) GetAlarms(ctx echo.Context, params server.GetAlarmsParams) error {
	return nil
}

func (b *BahnAlarmApi) DeleteAlarmsId(ctx echo.Context, id int) error {
	return nil
}
