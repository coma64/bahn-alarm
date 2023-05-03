package handlers

import (
	"context"
	"github.com/coma64/bahn-alarm-backend/server"
)

func (b *BahnAlarmApi) GetBahnConnections(ctx context.Context, request server.GetBahnConnectionsRequestObject) (server.GetBahnConnectionsResponseObject, error) {
	return server.GetBahnConnections200JSONResponse{}, nil
}

func (b *BahnAlarmApi) GetBahnPlaces(ctx context.Context, request server.GetBahnPlacesRequestObject) (server.GetBahnPlacesResponseObject, error) {
	return server.GetBahnPlaces200JSONResponse{}, nil
}
