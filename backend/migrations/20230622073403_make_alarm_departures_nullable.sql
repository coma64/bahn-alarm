-- +goose Up
-- +goose StatementBegin
alter table alarms alter column departureid drop not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
delete from alarms where departureId is null;
alter table alarms alter column departureId set not null;
-- +goose StatementEnd
