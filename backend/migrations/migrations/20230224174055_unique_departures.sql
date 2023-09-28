-- +goose Up
-- +goose StatementBegin
alter table departures add constraint departuresUniqueDeparture unique (connectionId, departure);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table departures drop constraint departuresUniqueDeparture;
-- +goose StatementEnd
