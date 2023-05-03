package handlers

import (
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
)

func (b *BahnAlarmApi) GetBahnConnections(ctx echo.Context, params server.GetBahnConnectionsParams) error {
	return nil
}

func (b *BahnAlarmApi) GetBahnPlaces(ctx echo.Context, params server.GetBahnPlacesParams) error {
	return nil
}
