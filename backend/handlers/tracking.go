package handlers

import (
	"context"
	"github.com/coma64/bahn-alarm-backend/server"
)

func (b *BahnAlarmApi) GetTrackingConnections(ctx context.Context, request server.GetTrackingConnectionsRequestObject) (server.GetTrackingConnectionsResponseObject, error) {
	return server.GetTrackingConnections200JSONResponse{}, nil
}

func (b *BahnAlarmApi) PostTrackingConnections(ctx context.Context, request server.PostTrackingConnectionsRequestObject) (server.PostTrackingConnectionsResponseObject, error) {
	return server.PostTrackingConnections201JSONResponse{}, nil
}

func (b *BahnAlarmApi) DeleteTrackingConnectionsId(ctx context.Context, request server.DeleteTrackingConnectionsIdRequestObject) (server.DeleteTrackingConnectionsIdResponseObject, error) {
	return server.DeleteTrackingConnectionsId204Response{}, nil
}

func (b *BahnAlarmApi) PutTrackingConnectionsId(ctx context.Context, request server.PutTrackingConnectionsIdRequestObject) (server.PutTrackingConnectionsIdResponseObject, error) {
	return server.PutTrackingConnectionsId204Response{}, nil
}

func (b *BahnAlarmApi) GetTrackingStats(ctx context.Context, request server.GetTrackingStatsRequestObject) (server.GetTrackingStatsResponseObject, error) {
	return server.GetTrackingStats200JSONResponse{}, nil
}
