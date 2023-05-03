package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type dbTrackedConnection struct {
	Id       int
	FromId   string
	FromName string
	ToId     string
	ToName   string
}

func extractConnectionIds(conns []dbTrackedConnection) []int {
	ids := make([]int, 0, len(conns))
	for _, conn := range conns {
		ids = append(ids, conn.Id)
	}
	return ids
}

func dbConnsToSchemaConns(conns []dbTrackedConnection) []server.TrackedConnection {
	schemas := make([]server.TrackedConnection, 0, len(conns))
	for i, conn := range conns {
		schemas = append(schemas, server.TrackedConnection{
			Departures: []server.TrackedDeparture{},
			From: server.BahnStation{
				Id:   conn.FromId,
				Name: conn.FromName,
			},
			Id: &conns[i].Id,
			To: server.BahnStation{
				Id:   conn.ToId,
				Name: conn.ToName,
			},
		})
	}
	return schemas
}

func (b *BahnAlarmApi) GetTrackingConnections(ctx echo.Context, params server.GetTrackingConnectionsParams) error {
	offset, size := defaultPagination(params.Page, params.Size)
	// TODO: maybe use normal scan instead?
	var conns []dbTrackedConnection
	if err := db.Db.SelectContext(
		ctx.Request().Context(),
		&conns,
		`
select
	c.id,
	f.externalId fromId,
	f.name fromName,
	t.externalId toId,
	t.name toName
from connections c
	inner join users u on c.trackedById = u.id
	inner join bahnStations f on c.fromId = f.id
	inner join bahnStations t on c.toId = t.id
where
	u.name = $1
offset $2 fetch first $3 rows only
		`,
		ctx.Get("username"),
		offset,
		size,
	); err != nil {
		return fmt.Errorf("error getting tracked conns: %w", err)
	}

	var pagination server.Pagination
	if err := db.Db.GetContext(
		ctx.Request().Context(),
		&pagination,
		"select count(c.id) totalItems from connections c inner join users u on c.trackedById = u.id where u.name = $1",
		ctx.Get("username"),
	); err != nil {
		return fmt.Errorf("error getting total items: %w", err)
	}

	if len(conns) == 0 {
		return ctx.JSON(http.StatusOK, server.TrackedConnectionList{
			Connections: dbConnsToSchemaConns(conns),
			Pagination:  pagination,
		})
	}

	departuresQuery, args, err := sqlx.In(`
select connectionId, departure, delayMinutes, status from fatDepartures where connectionId in (?)
		`,
		extractConnectionIds(conns),
	)
	if err != nil {
		return fmt.Errorf("error creating departures query: %w", err)
	}

	departuresQuery = db.Db.Rebind(departuresQuery)

	var departures []struct {
		server.TrackedDeparture
		DelayMinutes int
		ConnectionId int
	}
	if err = db.Db.SelectContext(
		ctx.Request().Context(),
		&departures,
		departuresQuery,
		args...,
	); err != nil {
		return fmt.Errorf("error getting connection departures: %w", err)
	}

	connSchemas := dbConnsToSchemaConns(conns)
	for _, departure := range departures {
		for connIndex := range connSchemas {
			if departure.ConnectionId == conns[connIndex].Id {
				departure.TrackedDeparture.Delay = departure.DelayMinutes
				connSchemas[connIndex].Departures = append(connSchemas[connIndex].Departures, departure.TrackedDeparture)
			}
		}
	}

	for _, conn := range connSchemas {
		sort.Slice(conn.Departures, func(i, j int) bool {
			return conn.Departures[i].Departure.Before(conn.Departures[j].Departure)
		})
	}

	return ctx.JSON(http.StatusOK, server.TrackedConnectionList{
		Connections: connSchemas,
		Pagination:  pagination,
	})
}

func (b *BahnAlarmApi) PostTrackingConnections(ctx echo.Context) error {
	var body server.PostTrackingConnectionsJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	tx, err := db.Db.BeginTxx(ctx.Request().Context(), nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	if _, err = tx.ExecContext(
		ctx.Request().Context(),
		"insert into bahnStations (externalId, name) values ($1, $2), ($3, $4) on conflict do nothing",
		body.From.Id,
		body.From.Name,
		body.To.Id,
		body.To.Name,
	); err != nil {
		return fmt.Errorf("error inserting bahn stations: %w", err)
	}

	response := &server.TrackedConnection{
		From:       body.From,
		To:         body.To,
		Departures: []server.TrackedDeparture{},
	}

	if err = tx.QueryRowxContext(
		ctx.Request().Context(),
		`
insert into connections
	(trackedById, fromId, toId)
values (
	(select id from users where name = $1),
	(select id from bahnStations where externalId = $2),
	(select id from bahnStations where externalId = $3)
) returning id
		`,
		ctx.Get("username"),
		body.From.Id,
		body.To.Id,
	).Scan(&response.Id); err != nil {
		_ = tx.Rollback()
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Constraint == "connections_trackedbyid_fromid_toid_key" {
			return ctx.NoContent(http.StatusConflict)
		}

		return fmt.Errorf("error inserting connection: %w", err)
	}

	if len(body.Departures) == 0 {
		if err = tx.Commit(); err != nil {
			return fmt.Errorf("error committing transaction: %w", err)
		}

		return ctx.JSON(http.StatusCreated, response)
	}

	query := db.Sq.Insert("departures").Columns("connectionId", "departure")
	for _, departure := range body.Departures {
		query = query.Values(response.Id, departure.Departure)
	}

	if _, err = query.RunWith(tx).ExecContext(ctx.Request().Context()); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("error inserting departures: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	for _, departure := range body.Departures {
		responseDeparture := server.TrackedDeparture{
			Departure: departure.Departure,
			Status:    server.NotChecked,
		}
		response.Departures = append(response.Departures, responseDeparture)
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (b *BahnAlarmApi) DeleteTrackingConnectionsId(ctx echo.Context, id int) error {
	result, err := db.Db.ExecContext(
		ctx.Request().Context(),
		"delete from connections c using users u where c.trackedById = u.id and u.name = $1 and c.id = $2",
		ctx.Get("username"),
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting connection: %w", err)
	}

	var affected int64
	if affected, err = result.RowsAffected(); err != nil {
		return fmt.Errorf("error getting affected rows: %w", err)
	}

	if affected == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusNoContent)
}

func extractDepartureTimes(departures []server.TrackedDepartureWrite) []time.Time {
	times := make([]time.Time, 0, len(departures))
	for _, d := range departures {
		times = append(times, d.Departure)
	}
	return times
}

func (b *BahnAlarmApi) PutTrackingConnectionsId(ctx echo.Context, id int) error {
	var body server.PutTrackingConnectionsIdJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}

	var hasMatchingRow int
	if err := db.Db.QueryRowxContext(
		ctx.Request().Context(),
		"select 1 from connections c inner join users u on c.trackedById = u.id where c.id = $1 and u.name = $2",
		id,
		ctx.Get("username"),
	).Scan(&hasMatchingRow); err != nil {
		if err == sql.ErrNoRows {
			return ctx.NoContent(http.StatusNotFound)
		}
		return fmt.Errorf("error checking for mathing connection: %w", err)
	}

	tx, err := db.Db.BeginTxx(ctx.Request().Context(), nil)
	if err != nil {
		return fmt.Errorf("error starting transactions: %w", err)
	}

	if _, err = db.Sq.Delete("departures").
		Where(squirrel.And{
			squirrel.Eq{"connectionId": id},
			squirrel.NotEq{"departure": extractDepartureTimes(body.Departures)},
		}).
		RunWith(tx).
		ExecContext(ctx.Request().Context()); err != nil {
		return fmt.Errorf("error deleting old departures: %w", err)
	}

	if len(body.Departures) == 0 {
		_ = tx.Rollback()
		return ctx.NoContent(http.StatusNoContent)
	}

	query := db.Sq.Insert("departures").Columns("connectionId", "departure").Suffix("on conflict do nothing")
	for _, departure := range body.Departures {
		query = query.Values(id, departure.Departure)
	}

	if _, err = query.RunWith(tx).ExecContext(ctx.Request().Context()); err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("error inserting added departures: %w", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (b *BahnAlarmApi) GetTrackingStats(ctx echo.Context) error {
	var stats server.TrackingStats
	if err := db.Db.GetContext(
		ctx.Request().Context(),
		&stats,
		`
select
    count(d.id) totalConnectionCount,
    count(d.id) filter (where i.actualTime is not null and i.actualTime != i.scheduledTime) delayedConnectionCount,
    count(d.id) filter (where i.departureId is not null and i.actualTime is null) canceledConnectionCount
from departures d
    left outer join departureInfos i on i.departureId = d.id
	inner join connections c on c.id = d.connectionId
	inner join users u on u.id = c.trackedById
where u.name = $1
		`,
		ctx.Get("username"),
	); err != nil {
		return fmt.Errorf("error getting connection counts: %w", err)
	}

	if err := db.Db.GetContext(
		ctx.Request().Context(),
		&stats.NextDeparture,
		`
select
    d.departure,
    d.connectionId
from departures d
    inner join connections c on c.id = d.connectionId
    inner join users u on u.id = c.trackedById
where u.name = $1
order by seconds_until_departure(d.departure)
fetch first 1 row only
		`,
		ctx.Get("username"),
	); err != nil {
		if err != sql.ErrNoRows {
			return fmt.Errorf("error getting next departure: %w", err)
		}
	}

	return ctx.JSON(http.StatusOK, &stats)
}
