package queries

import (
	"context"
	"database/sql"
	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/external_apis/bahn"
	"time"
)

type DepartureModel struct {
	models.Departure
	FromStationId          string
	FromStationName        string
	ToStationId            string
	ToStationName          string
	DepartureMarginMinutes int
	TrackedById            int
}

func TimeOnly(t time.Time) time.Time {
	year, month, day := t.Date()
	return t.AddDate(-year, -int(month)+1, -day+1)
}

func (d *DepartureModel) TimeUntilNextDeparture() time.Duration {
	nowTime := TimeOnly(time.Now().UTC())
	diff := d.Departure.Departure.Sub(nowTime)
	if d.Departure.Departure.Before(nowTime) {
		diff += time.Hour * 24
	}
	return diff
}

func SelectDueDepartures(ctx context.Context) ([]DepartureModel, error) {
	departures := []DepartureModel{}
	return departures, db.Db.SelectContext(
		ctx,
		&departures,
		`
select
	d.*,
	f.externalId fromStationId,
	f.name fromStationName,
	t.externalId toStationId,
	t.name toStationName,
	c.departureMarginMinutes,
	c.trackedById
from departures d
	inner join connections c on d.connectionId = c.id
	inner join bahnStations f on c.fromId = f.id
	inner join bahnStations t on c.toId = t.id
where d.nextCheck < now();
`,
	)
}

func UpdateDepartureNextCheck(ctx context.Context, id int, newNextCheck time.Time) error {
	_, err := db.Db.ExecContext(ctx, "update departures set nextCheck = $1 where id = $2", newNextCheck, id)
	return err
}

func CreateOrUpdateDepartureInfo(ctx context.Context, departure *DepartureModel, trip *bahn.Trip) (*models.DepartureInfo, error) {
	var departureInfo models.DepartureInfo
	return &departureInfo, db.Db.GetContext(
		ctx,
		&departureInfo,
		`
insert into departureInfos (departureId, scheduledTime, actualTime)
values ($1, $2, $3)
on conflict (departureid, day) do update set scheduledTime = $2, actualTime = $3
returning *
`,
		departure.Departure.Id,
		trip.Departure.ScheduledTime,
		trip.Departure.ActualTime,
	)
}

func GetDepartureInfos(ctx context.Context, departureId int) (*models.DepartureInfo, error) {
	var departure models.DepartureInfo
	if err := db.Db.GetContext(
		ctx,
		&departure,
		"select * from departureInfos where departureId = $1",
		departureId,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &departure, nil
}

func GetNextCheck(ctx context.Context) (*time.Time, error) {
	result := struct {
		NextCheck time.Time
	}{}

	return &result.NextCheck, db.Db.GetContext(ctx, &result, "select min(nextCheck) nextCheck from departures;")
}
