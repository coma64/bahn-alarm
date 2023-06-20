package queries

import (
	"context"
	"database/sql"
	"time"

	"github.com/coma64/bahn-alarm-backend/db"
	"github.com/coma64/bahn-alarm-backend/db/models"
	"github.com/coma64/bahn-alarm-backend/external_apis/bahn"
)

func SelectDueDepartures(ctx context.Context) ([]FatDeparture, error) {
	departures := []FatDeparture{}
	return departures, db.Db.SelectContext(
		ctx,
		&departures,
		`select * from fatDepartures where nextCheck < now()`,
	)
}

func UpdateDepartureNextCheck(ctx context.Context, id int, newNextCheck time.Time) error {
	_, err := db.Db.ExecContext(ctx, "update departures set nextCheck = $1 where id = $2", newNextCheck, id)
	return err
}

func CreateOrUpdateDepartureInfo(ctx context.Context, departure *FatDeparture, trip *bahn.Trip) (*models.DepartureInfo, error) {
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

	return &result.NextCheck, db.Db.GetContext(ctx, &result, "select coalesce(min(nextCheck), now() + '5 minutes') nextCheck from departures;")
}
