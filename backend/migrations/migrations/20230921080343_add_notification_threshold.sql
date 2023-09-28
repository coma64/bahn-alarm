-- +goose Up
-- +goose StatementBegin
alter table users add column notificationThresholdMinutes int default 5 not null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users drop column notificationThresholdMinutes;
-- +goose StatementEnd
