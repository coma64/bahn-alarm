package handlers

import (
	"context"
	"github.com/coma64/bahn-alarm-backend/server"
)

func (b *BahnAlarmApi) GetAlarms(ctx context.Context, request server.GetAlarmsRequestObject) (server.GetAlarmsResponseObject, error) {
	return server.GetAlarms200JSONResponse{}, nil
}

func (b *BahnAlarmApi) DeleteAlarmsId(ctx context.Context, request server.DeleteAlarmsIdRequestObject) (server.DeleteAlarmsIdResponseObject, error) {
	return server.DeleteAlarmsId204Response{}, nil
}
