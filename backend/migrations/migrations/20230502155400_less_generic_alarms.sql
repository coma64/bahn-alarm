-- +goose Up
-- +goose StatementBegin
delete from alarms;
alter table alarms
    drop column alarmdata,
    add column departureId bigint references departures not null,
    add column message text not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from alarms;
alter table alarms
    add column alarmdata jsonb not null,
    drop column departureId,
    drop column message;
-- +goose StatementEnd
