-- +goose Up
-- +goose StatementBegin
create type alarmUrgency as enum ('info', 'warn', 'error');
alter table alarms add column urgency alarmUrgency default 'info' not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop type alarmUrgency;
alter table alarms drop column urgency;
-- +goose StatementEnd
