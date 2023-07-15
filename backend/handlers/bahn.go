package handlers

import (
	"github.com/coma64/bahn-alarm-backend/external_apis/bahn"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func bahnConnectionToSchema(conn *bahn.Trip) server.BahnConnection {
	return server.BahnConnection{
		Departure: struct {
			ScheduledTime time.Time `json:"scheduledTime"`
		}{
			ScheduledTime: conn.Departure.ScheduledTime,
		},
	}
}

func (b *BahnAlarmApi) GetBahnConnections(ctx echo.Context, params server.GetBahnConnectionsParams) error {
	connections, err := bahn.FetchConnections(ctx.Request().Context(), params.Departure, params.FromId, params.ToId)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusServiceUnavailable,
			Message:  "cannot reach bahn API",
			Internal: err,
		}
	}

	schemas := make([]server.BahnConnection, 0, len(connections.Trips))
	for _, conn := range connections.Trips {
		schemas = append(schemas, bahnConnectionToSchema(&conn))
	}

	return ctx.JSON(http.StatusOK, server.BahnConnectionsList{Connections: schemas})
}

func bahnPlaceToSchema(place *bahn.Place) server.BahnStation {
	return server.BahnStation{
		Id:   place.StationID,
		Name: place.Name,
	}
}

func (b *BahnAlarmApi) GetBahnPlaces(ctx echo.Context, params server.GetBahnPlacesParams) error {
	if len(params.Name) < 2 {
		return ctx.NoContent(http.StatusBadRequest)
	}

	places, err := bahn.FetchPlaces(ctx.Request().Context(), params.Name)
	if err != nil {
		return &echo.HTTPError{
			Code:     http.StatusServiceUnavailable,
			Message:  "cannot reach bahn API",
			Internal: err,
		}
	}

	schemas := make([]server.BahnStation, 0, len(places.Places))
	for _, place := range places.Places {
		schemas = append(schemas, bahnPlaceToSchema(&place))
	}

	return ctx.JSON(http.StatusOK, server.BahnPlacesList{Places: schemas})
}
