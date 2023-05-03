package handlers

import (
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/server"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (b *BahnAlarmApi) GetAlarms(ctx echo.Context, params server.GetAlarmsParams) error {
	offset, size := defaultPagination(params.Page, params.Size)
	isNotFilteringUrgency := params.Urgency == nil

	alarms := []struct {
		models.Alarm
		server.SimpleConnection
	}{}
	if err := db.Db.SelectContext(
		ctx.Request().Context(),
		&alarms,
		`
select
    a.*,
    f.departure,
    f.fromStationName fromName,
    f.toStationName toName
from alarms a
    inner join users u on a.receiverId = u.id
	inner join fatDepartures f on f.id = a.departureId
where u.name = $1 and ($2 or a.urgency = $3)
order by a.createdAt desc
offset $4 fetch first $5 rows only
`,
		ctx.Get("username"),
		isNotFilteringUrgency,
		params.Urgency,
		offset,
		size,
	); err != nil {
		return err
	}

	var pagination server.Pagination
	if err := db.Db.GetContext(
		ctx.Request().Context(),
		&pagination,
		"select count(a.id) totalItems from alarms a inner join users u on a.receiverId = u.id where u.name = $1 and ($2 or a.urgency = $3)",
		ctx.Get("username"),
		isNotFilteringUrgency,
		params.Urgency,
	); err != nil {
		return err
	}

	alarmSchemas := make([]server.Alarm, 0, len(alarms))
	for _, alarm := range alarms {
		schema, err := alarm.ToSchema(&alarm.SimpleConnection)
		if err != nil {
			return err
		}

		alarmSchemas = append(alarmSchemas, *schema)
	}

	return ctx.JSON(
		http.StatusOK,
		server.AlarmsList{
			Alarms:     alarmSchemas,
			Pagination: pagination,
		},
	)
}

func (b *BahnAlarmApi) DeleteAlarmsId(ctx echo.Context, id int) error {
	result, err := db.Db.ExecContext(
		ctx.Request().Context(),
		"delete from alarms a using users u where a.receiverId = u.id and u.name = $1 and a.id = $2",
		ctx.Get("username"),
		id,
	)
	if err != nil {
		return err
	}

	var rows int64
	if rows, err = result.RowsAffected(); err != nil {
		return err
	}

	if rows == 0 {
		return ctx.NoContent(http.StatusNotFound)
	}
	return ctx.NoContent(http.StatusNoContent)
}
